import useReactWebSocket, { ReadyState } from "react-use-websocket";
import { getAPIURL } from "../../util/fetch";
import { useAuth0 } from "@auth0/auth0-react";
import {
  ClientMessage,
  ClientMessageType,
  ServerMessage,
  ServerMessageType,
} from "../../proto/websocket/v1/websocket";
import {
  createContext,
  createElement,
  PropsWithChildren,
  useContext,
  useEffect,
  useMemo,
  useRef,
} from "react";

type WebsocketProviderProps = Record<string, unknown>;

export type EventListener = (message: ServerMessage) => void;
type AttachEventListener = (
  type: ServerMessageType,
  listener: EventListener
) => void;

type ContextValue = {
  addEventListener: AttachEventListener;
  removeEventListener: AttachEventListener;
};

const noop: AttachEventListener = () => {};

const WebsocketContext = createContext<ContextValue>({
  addEventListener: noop,
  removeEventListener: noop,
});

const makeMessage = <P>(type: ClientMessageType, payload: P): ClientMessage => {
  return {
    timestamp: new Date(),
    type,
    payloadVersion: "v1",
    ...payload,
  };
};

const makeTokenMessage = (token: string): ClientMessage => {
  return makeMessage(ClientMessageType.CLIENT_MESSAGE_TYPE_TOKEN_UPDATE, {
    tokenUpdate: {
      token,
    },
  });
};

export const WebsocketProvider = ({
  children,
}: PropsWithChildren<WebsocketProviderProps>) => {
  const { getAccessTokenSilently } = useAuth0();
  const url = `${getAPIURL()}/ws`;

  const { sendJsonMessage, lastJsonMessage, readyState } = useReactWebSocket(
    url,
    {
      onOpen: async () => {
        let token = "";
        try {
          token = await getAccessTokenSilently();
        } catch (e) {
          //showError("Unauthorized");
          return;
        }

        sendJsonMessage(makeTokenMessage(token));
      },
      retryOnError: true,
      shouldReconnect: () => true,
    }
  );

  useEffect(() => {
    const timer = setInterval(async () => {
      if (readyState === ReadyState.OPEN) {
        let token = "";
        try {
          token = await getAccessTokenSilently();
        } catch (e) {
          //showError("Unauthorized");
          return;
        }

        sendJsonMessage(makeTokenMessage(token));
      }
    }, 10000);

    return () => clearInterval(timer);
  }, [readyState, sendJsonMessage]);

  const listeners = useRef(new Map<ServerMessageType, EventListener[]>());

  const value = useMemo<ContextValue>(() => {
    return {
      addEventListener: (type, cb) => {
        if (!listeners.current.has(type)) {
          listeners.current.set(type, []);
        }

        listeners.current.get(type)?.push(cb);
      },
      removeEventListener: (type, cb) => {
        if (!listeners.current.has(type)) {
          return;
        }

        listeners.current.set(
          type,
          listeners.current.get(type)?.filter((savedCB) => savedCB !== cb) ?? []
        );
      },
    };
  }, [listeners]);

  useEffect(() => {
    const message = lastJsonMessage as ServerMessage;
    if (lastJsonMessage) {
      const enumKey = message.type as unknown as keyof typeof ServerMessageType;
      listeners.current.get(ServerMessageType[enumKey])?.forEach((listener) => {
        listener(message);
      });
    }
  }, [lastJsonMessage, listeners]);

  return createElement(WebsocketContext.Provider, { value: value }, children);
};

export const useWebsocketContext = () => useContext(WebsocketContext);
