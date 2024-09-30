package service

import (
	"context"

	"api/internal/api"
	"api/internal/domain"
	"api/pkg/syserr"
)

// StoreAPIService is a service that implements the logic for the StoreAPIServicer
// This service should implement the business logic for every endpoint for the StoreAPI API.
// Include any external packages or services that will be required by this service.
type StoreAPIService struct {
}

// NewStoreAPIService creates a default api service
func NewStoreAPIService() *StoreAPIService {
	return &StoreAPIService{}
}

// DeleteOrder - Delete purchase order by ID
func (s *StoreAPIService) DeleteOrder(ctx context.Context, orderId int64) (*domain.DeleteOrderResponse, error) {
	return nil, syserr.NewNotImplemented("")
}

// GetInventory - Returns pet inventories by status
func (s *StoreAPIService) GetInventory(ctx context.Context) (*domain.GetInventoryResponse, error) {
	return nil, syserr.NewNotImplemented("")
}

// GetOrderById - Find purchase order by ID
func (s *StoreAPIService) GetOrderById(ctx context.Context, orderId int64) (*domain.GetOrderByIdResponse, error) {
	return nil, syserr.NewNotImplemented("")
}

// PlaceOrder - Place an order for a pet
func (s *StoreAPIService) PlaceOrder(ctx context.Context, order api.Order) (*domain.PlaceOrderResponse, error) {
	return nil, syserr.NewNotImplemented("")
}
