package orders

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	ID              int             `json:"id"`
	OrderNumber     string          `json:"order_number"`
	Status          string          `json:"status"`
	ShippingAddress string          `json:"shipping_address"`
	UserID          int             `json:"user_id"`
	Total           decimal.Decimal `json:"total"`
	TotalDiscount   decimal.Decimal `json:"total_discount"`
	IsCancelled     bool            `json:"is_cancelled"`
	OrderItems      []OrderItem     `json:"order_items"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type OrderItem struct {
	ID        int             `json:"id"`
	OrderID   int             `json:"order_id"`
	ProductID string          `json:"product_id"`
	Quantity  int             `json:"quantity"`
	Price     decimal.Decimal `json:"price"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
