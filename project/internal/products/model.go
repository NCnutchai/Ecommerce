package products

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID          int             `json:"id"`
	Code        string          `json:"code"`
	ProductName string          `json:"product"`
	Price       decimal.Decimal `json:"price"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}
