package factory

import (
	"worker/internal/interfaces"

	"go.temporal.io/sdk/client"
)

type Factory struct {
	temporalClient client.Client
	temporalService interfaces.TemporalService
}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) SetTemporalClient(tempClient client.Client) {
	f.temporalClient = tempClient
}
