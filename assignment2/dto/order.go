package dto

import (
	"assignment-2/entity"
	"time"
)

type NewOrderRequest struct {
	OrderedAt    time.Time        `json:"orderedAt"`
	CustomerName string           `json:"customerName"`
	Items        []NewItemRequest `json:"items"`
}

func (o *NewOrderRequest) OrderRequestToEntity() *entity.Order {
	return &entity.Order{
		CustomerName: o.CustomerName,
		OrderedAt:    o.OrderedAt,
	}
}

type NewOrderResponse struct {
	StatusCode int             `json:"statusCode"`
	Message    string          `json:"message"`
	Data       NewOrderRequest `json:"data"`
}

type GetAllOrdersResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       []OrderData `json:"data"`
}

type GetOrderByIDResponse struct {
	StatusCode int       `json:"statusCode"`
	Message    string    `json:"message"`
	Data       OrderData `json:"data"`
}

type DeleteOrderByIDResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type OrderData struct {
	ID           uint       `json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	CustomerName string     `json:"customerName"`
	Items        []ItemData `json:"items"`
}
