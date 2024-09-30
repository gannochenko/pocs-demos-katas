package service

import (
	"context"

	"api/internal/api"
	"api/internal/domain"
	"api/pkg/syserr"
)

// StoreService is a service that implements the logic for the StoreAPIServicer
// This service should implement the business logic for every endpoint for the StoreAPI API.
// Include any external packages or services that will be required by this service.
type StoreService struct {
}

// NewStoreService creates a default api service
func NewStoreService() *StoreService {
	return &StoreService{}
}

// DeleteOrder - Delete purchase order by ID
func (s *StoreService) DeleteOrder(ctx context.Context, orderId int64) (*domain.DeleteOrderResponse, error) {
	return nil, syserr.NewNotImplemented("")
}

// GetInventory - Returns pet inventories by status
func (s *StoreService) GetInventory(ctx context.Context) (*domain.GetInventoryResponse, error) {
	return nil, syserr.NewNotImplemented("")
}

// GetOrderById - Find purchase order by ID
func (s *StoreService) GetOrderById(ctx context.Context, orderId int64) (*domain.GetOrderByIdResponse, error) {
	return nil, syserr.NewNotImplemented("")
}

// PlaceOrder - Place an order for a pet
func (s *StoreService) PlaceOrder(ctx context.Context, order api.Order) (*domain.PlaceOrderResponse, error) {
	return nil, syserr.NewNotImplemented("")
}
