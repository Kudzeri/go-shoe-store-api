package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	OrderNumber     string      `json:"order_number" gorm:"type:varchar(100);not null;unique"`
	UserID          uint        `json:"user_id" gorm:"not null"`
	TotalAmount     float64     `json:"total_amount" gorm:"type:decimal(10,2);not null"`
	OrderStatus     string      `json:"order_status" gorm:"type:varchar(50);not null"`
	OrderDate       time.Time   `json:"order_date" gorm:"not null"`
	DeliveryDate    *time.Time  `json:"delivery_date"`
	ShippingAddress string      `json:"shipping_address" gorm:"type:text;not null"`
	Items           []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id" gorm:"not null"`
	ProductID uint    `json:"product_id" gorm:"not null"`
	Quantity  int     `json:"quantity" gorm:"not null"`
	UnitPrice float64 `json:"unit_price" gorm:"type:decimal(10,2);not null"`
}
