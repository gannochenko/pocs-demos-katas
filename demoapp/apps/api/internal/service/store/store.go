package store

import (
	"context"

	"api/internal/api"
	"api/internal/domain"
	"api/pkg/syserr"
)

// Service is a service that implements the logic for the StoreAPIServicer
// This service should implement the business logic for every endpoint for the StoreAPI API.
// Include any external packages or services that will be required by this service.
type Service struct {
}

// NewStoreService creates a default api service
func NewStoreService() *Service {
	return &Service{}
}

// DeleteOrder - Delete purchase order by ID
func (s *Service) DeleteOrder(ctx context.Context, orderId int64) (*domain.DeleteOrderResponse, error) {
	return nil, syserr.NewNotImplemented("method not implemented")
}

// GetInventory - Returns pet inventories by status
func (s *Service) GetInventory(ctx context.Context) (*domain.GetInventoryResponse, error) {
	return nil, syserr.NewNotImplemented("method not implemented")
}

// GetOrderById - Find purchase order by ID
func (s *Service) GetOrderById(ctx context.Context, orderId int64) (*domain.GetOrderByIdResponse, error) {
	return nil, syserr.NewNotImplemented("method not implemented")
}

// PlaceOrder - Place an order for a pet
func (s *Service) PlaceOrder(ctx context.Context, order api.Order) (*domain.PlaceOrderResponse, error) {
	return nil, syserr.NewNotImplemented("method not implemented")
}
