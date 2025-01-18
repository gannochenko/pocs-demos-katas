package network

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/samber/lo"
	"google.golang.org/protobuf/encoding/protojson"

	"backend/interfaces"
	ctxUtil "backend/internal/util/ctx"
	"backend/internal/util/syserr"
	"backend/internal/util/types"
	v1 "backend/proto/websocket/v1"
)

const (
	TOKEN_PROVISION_TIMEOUT = 3
)

type WebsocketConnection struct {
	id           uuid.UUID
	userID       *uuid.UUID
	outgoingChan chan *v1.ServerMessage
	incomingChan chan incomingMessage
	expiresAt    *time.Time
	wsConnection *websocket.Conn
}

func (c *WebsocketConnection) IsValid() bool {
	return c.expiresAt.After(time.Now())
}

func (c *WebsocketConnection) NotValid() <-chan struct{} {
	done := make(chan struct{})

	if c.expiresAt == nil {
		close(done)
		return done
	}

	duration := time.Until(*c.expiresAt)
	if duration <= 0 {
		close(done)
		return done
	}

	go func() {
		time.Sleep(duration)
		close(done)
	}()

	return done
}

func (c *WebsocketConnection) Close() error {
	close(c.outgoingChan)
	close(c.incomingChan)

	err := c.wsConnection.Close()
	if err != nil {
		return syserr.Wrap(err, "could not close websocket connection")
	}

	return nil
}

type incomingMessage struct {
	messageType int
	message     []byte
	err         error
}

type WebsocketServer struct {
	configService interfaces.ConfigService
	authService   interfaces.AuthService
	loggerService interfaces.LoggerService
	userService   interfaces.UserService
	connections   sync.Map
}

func NewWebsocketServer(
	configService interfaces.ConfigService,
	authService interfaces.AuthService,
	loggerService interfaces.LoggerService,
	userService interfaces.UserService,
) *WebsocketServer {
	return &WebsocketServer{
		configService: configService,
		authService:   authService,
		loggerService: loggerService,
		userService:   userService,
	}
}

func (s *WebsocketServer) Start(ctx context.Context) error {
	return nil
}

func (s *WebsocketServer) Stop() error {
	return nil
}

func (s *WebsocketServer) GetHandler() types.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		// todo: record error with defer, if needed

		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		upgrader, err := s.getUpgrader()
		if err != nil {
			return syserr.Wrap(err, "could not get upgrader")
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return syserr.Wrap(err, "could not upgrade connection")
		}

		connection := &WebsocketConnection{
			id:           uuid.New(),
			userID:       nil, // user will be assigned later
			outgoingChan: make(chan *v1.ServerMessage),
			incomingChan: s.createIncomingChannel(ctx, conn, cancel),
			expiresAt:    nil, // expiration will be assigned later
			wsConnection: conn,
		}

		defer func() {
			err := s.closeAndRemoveConnection(connection)
			if err != nil {
				s.loggerService.LogError(ctx, syserr.Wrap(err, "could not close and remove connection"))
			} else {
				s.loggerService.Info(ctx, "connection closed and removed")
			}
		}()

		s.loggerService.Info(ctx, "websocket is waiting for the token")

		err = s.runHandshake(ctx, connection)
		if err != nil {
			return syserr.Wrap(err, "could not conduct handshake")
		}

		//ctx = ctxUtil.WithUserEmail(ctx, userEmail)

		connection, err := s.addConnection(ctx, userEmail, connectionID, &authorizationExpiresAt)
		if err != nil {
			return syserr.WrapAsInternal(err, "error adding a connection")
		}

		s.loggerService.Info(ctx, "started listening to messages")

		for {
			select {
			case <-ctx.Done():
				return syserr.NewInternal("context is done, closing the connection")
			case <-connection.NotValid():
				s.loggerService.Warning(ctx, "the token was not updated in time, closing the connection")
				return nil
			case message := <-connection.outgoingChannel:
				data, err := protojsonOptions.Marshal(message)
				if err != nil {

					log.SmartLog(ctx, syserr.WrapAsInternal(err, "error marshaling an outgoing message"))
					continue
				}

				err = conn.WriteMessage(websocket.TextMessage, data)
				if err != nil {
					s.loggerService.LogError(ctx, syserr.Wrap(err, "error sending an outgoing message"))
				}
			case message := <-incomingChannel:
				switch message.messageType {
				case websocket.TextMessage:
					err = s.processMessage(ctx, userEmail, connectionID, message.message)
					if err != nil {
						s.loggerService.LogError(ctx, syserr.Wrap(err, "error processing an incoming message"))
					}
				case websocket.CloseMessage:
					s.loggerService.Info(ctx, "close message received, closing the connection")
					return nil
				}
			}
		}
	}
}

func (s *WebsocketServer) runHandshake(ctx context.Context, connection *WebsocketConnection) error {
	for {
		select {
		case <-ctx.Done():
			return syserr.Wrap(context.Canceled, "ctx done, exiting")
		case message := <-connection.incomingChan:
			s.loggerService.Info(ctx, "token message received")

			protoMessage, err := s.unmarshalMessage(&message)
			if err != nil {
				return syserr.WrapAs(err, syserr.BadInputCode, "could not decode message")
			}

			if protoMessage.GetType() != v1.ClientMessageType_CLIENT_MESSAGE_TYPE_TOKEN_UPDATE {
				return syserr.WrapAs(err, syserr.BadInputCode, "first message is not of token update, closing the connection")
			}

			token := protoMessage.GetTokenUpdate().GetToken()

			sup, err := s.authService.ValidateToken(ctx, token)
			if err != nil {
				return syserr.NewBadInput("could not validate token")
			}

			user, err := s.userService.GetUserBySUP(ctx, nil, sup)
			if err != nil {
				return syserr.Wrap(err, "could not get user by sup")
			}

			connection.userID = &user.ID
			connection.expiresAt = info.ExpiresAt

			break
		case <-time.After(time.Second * TOKEN_PROVISION_TIMEOUT):
			return syserr.NewBadInput("no token provided before the timeout, closing the connection")
		}
	}
}

func (s *WebsocketServer) getUpgrader() (*websocket.Upgrader, error) {
	config, err := s.configService.GetConfig()
	if err != nil {
		return nil, syserr.Wrap(err, "could not get config")
	}

	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return lo.Contains(config.HTTP.Cors.Origin, r.Header.Get("Origin"))
		},
	}, nil
}

func (s *Service) addConnection(ctx context.Context, userEmail string, connectionID string, expiresAt *time.Time) (*Connection, error) {
	key := s.makeConnMapKey(userEmail, connectionID)
	if s.connections.HasKey(key) {
		return nil, syserror.NewInternal("connection already registered", syserror.F("user_email", userEmail), syserror.F("connection_id", connectionID))
	}

	connection := &Connection{
		outgoingChannel: make(chan *protopbV1.ServerMessage),
		expiresAt:       expiresAt,
	}

	s.connections.Set(key, connection)

	_, newValue := s.poolSize.Change(func(value uint32) uint32 {
		return value + 1
	})

	s.recordConnectionPoolSize(ctx, newValue)

	return connection, nil
}

func (s *WebsocketServer) closeAndRemoveConnection(connection *WebsocketConnection) error {
	err := connection.Close()
	if err != nil {
		return syserr.Wrap(err, "could not close connection")
	}

	if connection.userID != nil {
		// the connection has already been stored in the pool, remove it
		// todo: remove
	}

	hadConnection := s.hasConnection(userEmail, connectionID)

	key := s.makeConnMapKey(userEmail, connectionID)
	s.connections.Act(func(innerMap map[string]*Connection) {
		if maps.HasKey(innerMap, key) {
			close(innerMap[key].outgoingChannel)
			delete(innerMap, key)
		}
	})
}

func (s *Service) doesConnectionExistAndValid(userEmail string, connectionID string) bool {
	connection, ok := s.connections.Get(s.makeConnMapKey(userEmail, connectionID))
	if ok {
		return connection.IsValid()
	}

	return false
}

func (s *Service) hasConnection(userEmail string, connectionID string) bool {
	_, ok := s.connections.Get(s.makeConnMapKey(userEmail, connectionID))
	return ok
}

func (s *WebsocketServer) unmarshalMessage(message *incomingMessage) (*v1.ClientMessage, error) {
	var protoMessage v1.ClientMessage
	err := protojson.Unmarshal(message.message, &protoMessage)
	if err != nil {
		return nil, syserr.WrapAs(err, syserr.BadInputCode, "could not unmarshal message")
	}

	return &protoMessage, nil
}

func (s *WebsocketServer) createIncomingChannel(ctx context.Context, conn *websocket.Conn, cancel context.CancelFunc) chan incomingMessage {
	messageChan := make(chan incomingMessage)

	go func() {
		for {
			if ctxUtil.IsDone(ctx) {
				close(messageChan)
				return
			}

			messageType, message, err := conn.ReadMessage()
			if err != nil {
				s.loggerService.LogError(ctx, syserr.Wrap(err, "could not read incoming message"))
				close(messageChan)
				cancel()
				return
			}

			messageChan <- incomingMessage{
				messageType: messageType,
				message:     message,
			}
		}
	}()

	return messageChan
}

func (s *Service) processMessage(ctx context.Context, userEmail string, connectionID string, payload []byte) error {
	if !s.doesConnectionExistAndValid(userEmail, connectionID) {
		return syserror.NewInternal("connection was not valid when processing a message")
	}

	var protoMessage protopbV1.ClientMessage
	err := protojson.Unmarshal(payload, &protoMessage)
	if err != nil {
		return err
	}

	switch protoMessage.GetType() {
	case protopbV1.ClientMessageType_CLIENT_MESSAGE_TYPE_TOKEN_UPDATE:
		err = s.processTokenUpdateMessage(ctx, userEmail, connectionID, &protoMessage)
	default:
		err = syserror.NewBadInput("unrecognised message, skipped", syserror.F("payload", string(payload)))
	}

	return err
}

func (s *Service) processTokenUpdateMessage(ctx context.Context, userEmail string, connectionID string, protoMessage *protopbV1.ClientMessage) error {
	info, err := s.auth.ExtractAuthInfoFromToken(ctx, protoMessage.GetTokenUpdate().Token)
	if err != nil {
		return syserror.WrapAsInternal(err, "could not read the token")
	}

	conn, _ := s.connections.Get(s.makeConnMapKey(userEmail, connectionID))
	conn.expiresAt = &info.ExpiresAt

	log.InfoCtx(ctx, "token message was processed")

	return nil
}
