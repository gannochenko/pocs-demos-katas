package network

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/samber/lo"
	"google.golang.org/protobuf/encoding/protojson"

	"backend/interfaces"
	"backend/internal/domain"
	ctxUtil "backend/internal/util/ctx"
	"backend/internal/util/logger"
	"backend/internal/util/syserr"
	"backend/internal/util/types"
	v1 "backend/proto/websocket/v1"
)

const (
	TokenProvisionTimeout = 3
)

type WebsocketConnection struct {
	id           uuid.UUID
	userID       *uuid.UUID
	userSup      *string
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
	types.CloseChannelSafely(c.outgoingChan)
	types.CloseChannelSafely(c.incomingChan)

	err := c.wsConnection.Close()
	if err != nil {
		return syserr.Wrap(err, "could not close websocket connection")
	}

	return nil
}

func (c *WebsocketConnection) GetKey() string {
	return fmt.Sprintf("%s-%s", c.userID.String(), c.id.String())
}

type incomingMessage struct {
	messageType int
	message     []byte
	err         error
}

type WebsocketServer struct {
	configService   interfaces.ConfigService
	authService     interfaces.AuthService
	loggerService   interfaces.LoggerService
	userService     interfaces.UserService
	eventBusService interfaces.EventBusService
	connections     sync.Map
}

func NewWebsocketServer(
	configService interfaces.ConfigService,
	authService interfaces.AuthService,
	loggerService interfaces.LoggerService,
	userService interfaces.UserService,
	eventBusService interfaces.EventBusService,
) *WebsocketServer {
	return &WebsocketServer{
		configService:   configService,
		authService:     authService,
		loggerService:   loggerService,
		userService:     userService,
		eventBusService: eventBusService,
	}
}

func (s *WebsocketServer) Start(ctx context.Context) error {
	callback := func(event *domain.EventBusEvent) {
		s.loggerService.Info(ctx, "websocket: new event received", logger.F("event", event))
		err := s.DispatchMessage(event)
		if err != nil {
			s.loggerService.LogError(ctx, syserr.Wrap(err, "could not dispatch event"))
		}
	}

	err := s.eventBusService.AddEventListener(domain.EventBusEventTypeImageProcessed, callback)
	if err != nil {
		return syserr.Wrap(err, "could not start listening to events")
	}

	select {
	case <-ctx.Done():
		return nil
	}
}

func (s *WebsocketServer) Stop() error {
	// todo: call s.eventBusService.RemoveEventListener()
	return nil
}

func (s *WebsocketServer) DispatchMessage(event *domain.EventBusEvent) error {
	s.connections.Range(func(key, value interface{}) bool {
		// connection, ok := value.(*WebsocketConnection)
		// if !ok {
		// 	return true
		// }

		// connection.outgoingChan <- &v1.ServerMessage{
		// 	Type: v1.ServerMessageType_SERVER_MESSAGE_TYPE_EVENT,
		// 	Event: event,
		// }
		return true
	})
	
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

		user, err := s.runHandshake(ctx, connection)
		if err != nil {
			return syserr.Wrap(err, "could not conduct handshake")
		}

		ctx = ctxUtil.WithUser(ctx, *user)

		err = s.addConnection(connection)
		if err != nil {
			return syserr.Wrap(err, "error adding a connection")
		}

		return s.serveConnection(ctx, connection)
	}
}

func (s *WebsocketServer) serveConnection(ctx context.Context, connection *WebsocketConnection) error {
	s.loggerService.Info(ctx, "started listening to messages")

	for {
		select {
		case <-ctx.Done():
			return syserr.Wrap(context.Canceled, "context is done, closing the connection")
		case <-connection.NotValid():
			s.loggerService.Warning(ctx, "the token was not updated in time, closing the connection")
			return nil
		case message := <-connection.outgoingChan:
			data, err := protojson.MarshalOptions{
				EmitUnpopulated: true,
			}.Marshal(message)
			if err != nil {
				s.loggerService.LogError(ctx, syserr.Wrap(err, "error marshaling a message"))
				continue
			}

			err = connection.wsConnection.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				s.loggerService.LogError(ctx, syserr.Wrap(err, "error sending a message"))
			}
		case message := <-connection.incomingChan:
			switch message.messageType {
			case websocket.TextMessage:
				err := s.processMessage(ctx, connection, message.message)
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

func (s *WebsocketServer) runHandshake(ctx context.Context, connection *WebsocketConnection) (*domain.User, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, syserr.Wrap(context.Canceled, "ctx done, exiting")
		case message := <-connection.incomingChan:
			s.loggerService.Info(ctx, "token message received")

			protoMessage, err := s.unmarshalMessage(&message)
			if err != nil {
				return nil, syserr.WrapAs(err, syserr.BadInputCode, "could not decode message")
			}

			if protoMessage.GetType() != v1.ClientMessageType_CLIENT_MESSAGE_TYPE_TOKEN_UPDATE {
				return nil, syserr.WrapAs(err, syserr.BadInputCode, "first message is not of token update, closing the connection")
			}

			token := protoMessage.GetTokenUpdate().GetToken()

			sup, expiry, err := s.authService.ValidateToken(ctx, token)
			if err != nil {
				return nil, syserr.NewBadInput("could not validate token")
			}

			user, err := s.userService.GetUserBySUP(ctx, nil, sup)
			if err != nil {
				return nil, syserr.Wrap(err, "could not get user by sup")
			}

			connection.userID = &user.ID
			connection.userSup = &user.Sup
			connection.expiresAt = lo.ToPtr(time.Unix(expiry, 0))

			return user, nil
		case <-time.After(time.Second * TokenProvisionTimeout):
			return nil, syserr.NewBadInput("no token provided before the timeout, closing the connection")
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
			allowedOrigins := config.Backend.HTTP.Cors.Origin
			if lo.Contains(allowedOrigins, "*") {
				return true
			}

			return lo.Contains(allowedOrigins, r.Header.Get("Origin"))
		},
	}, nil
}

func (s *WebsocketServer) addConnection(connection *WebsocketConnection) error {
	key := connection.GetKey()
	_, hasKey := s.connections.LoadOrStore(key, connection)

	if hasKey {
		return syserr.NewInternal("connection already registered", syserr.F("user_id", connection.userID))
	}

	return nil
}

func (s *WebsocketServer) closeAndRemoveConnection(connection *WebsocketConnection) error {
	err := connection.Close()
	if err != nil {
		return syserr.Wrap(err, "could not close connection")
	}

	if connection.userID != nil {
		key := connection.GetKey()
		s.connections.Delete(key)
	}

	return nil
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
				types.CloseChannelSafely(messageChan)
				return
			}

			messageType, message, err := conn.ReadMessage()
			if err != nil {
				s.loggerService.LogError(ctx, syserr.Wrap(err, "could not read incoming message"))
				types.CloseChannelSafely(messageChan)
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

func (s *WebsocketServer) processMessage(ctx context.Context, connection *WebsocketConnection, payload []byte) error {
	var protoMessage v1.ClientMessage
	err := protojson.Unmarshal(payload, &protoMessage)
	if err != nil {
		return err
	}

	switch protoMessage.GetType() {
	case v1.ClientMessageType_CLIENT_MESSAGE_TYPE_TOKEN_UPDATE:
		err = s.processTokenUpdateMessage(ctx, connection, &protoMessage)
	// todo: add more types here when needed
	default:
		err = syserr.NewBadInput("unrecognised message, skipped", syserr.F("payload", string(payload)))
	}

	return err
}

func (s *WebsocketServer) processTokenUpdateMessage(ctx context.Context, connection *WebsocketConnection, protoMessage *v1.ClientMessage) error {
	token := protoMessage.GetTokenUpdate().GetToken()

	sup, expiry, err := s.authService.ValidateToken(ctx, token)
	if err != nil {
		return syserr.NewBadInput("could not validate token")
	}

	if *connection.userSup != sup {
		return syserr.NewBadInput("user in the token is different from the connection owner")
	}

	connection.expiresAt = lo.ToPtr(time.Unix(expiry, 0))

	s.loggerService.Info(ctx, "token message was processed")

	return nil
}
