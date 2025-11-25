package factory

import (
	"api/internal/interfaces"
	"api/internal/service/temporal"

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

func (f *Factory) GetTemporalService() interfaces.TemporalService {
	if f.temporalService == nil {
		f.temporalService = temporal.NewService(f.temporalClient)
	}
	return f.temporalService
}
