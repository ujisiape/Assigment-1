package models

import (
	"time"
)

type Orders struct {
	Order_ID     uint      `gorm:"primaryKey;column:order_id" json:"OrderID"`
	CustomerName string    `gorm:"type:VARCHAR(50);not null" json:"customerName"`
	Items        []Items   `json:"Items" gorm:"foreignKey:order_id;constraint:OnDelete:CASCADE"`
	OrderedAt    time.Time `gorm:"not null;type:timestamp;autoCreateTime" json:"orderedAt"`
}
type Items struct {
	Item_ID     uint   `gorm:"primaryKey;column:item_id" json:"lineItemId"`
	ItemCode    string `gorm:"not null;type:VARCHAR(20)" json:"itemCode"`
	Description string `gorm:"type:TEXT" json:"description" `
	Quantity    uint   `gorm:"not null" json:"quantity"`
	OrderID     uint   `gorm:"not null" json:"OrderID"`
}
type Response struct {
	Message string `json:"message"`
	Data    Orders `json:"data"`
}
