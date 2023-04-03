package dto

import (
	"assignment-2/entity"
	"time"
)

type NewItemRequest struct {
	ItemCode    string `json:"itemCode"    binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    uint   `json:"quantity"    binding:"required"`
}

func (i *NewItemRequest) ItemRequestToEntity() *entity.Item {
	return &entity.Item{
		ItemCode:    i.ItemCode,
		Description: i.Description,
		Quantity:    i.Quantity,
	}
}

type ItemData struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	ItemCode    string    `json:"itemCode"`
	Description string    `json:"description"`
	Quantity    uint      `json:"quantity"`
	OrderID     uint      `json:"orderId"`
}

type UpdateItemResponse struct {
	StatusCode int      `json:"statusCode"`
	Message    string   `json:"message"`
	Data       ItemData `json:"data"`
}
