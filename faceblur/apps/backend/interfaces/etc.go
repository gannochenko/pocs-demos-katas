package interfaces

import "backend/internal/util/types"

type WebsocketServer interface {
	GetHandler() types.HTTPHandler
}
