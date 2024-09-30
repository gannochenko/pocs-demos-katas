package domain

import "time"

type OrderStatus string

const (
	OrderStatusPlaced    OrderStatus = "placed"
	OrderStatusApproved  OrderStatus = "approved"
	OrderStatusDelivered OrderStatus = "delivered"
)

type Order struct {
	ID       int64       `json:"id,omitempty"`
	PetID    int64       `json:"petId,omitempty"`
	Quantity int32       `json:"quantity,omitempty"`
	ShipDate time.Time   `json:"shipDate,omitempty"`
	Status   OrderStatus `json:"status,omitempty"`
	Complete bool        `json:"complete,omitempty"`
}

type Customer struct {
	ID       int64     `json:"id,omitempty"`
	Username string    `json:"username,omitempty"`
	Address  []Address `json:"address,omitempty"`
}

type DeleteOrderResponse struct{}

type GetInventoryResponse struct{}

type GetOrderByIdResponse struct{}

type PlaceOrderResponse struct{}
