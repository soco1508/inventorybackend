package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Sale struct {
	SaleID      string          `db:"sale_id" json:"saleId"`
	ProductID   string          `db:"product_id" json:"productId"`
	Timestamp   time.Time       `db:"timestamp" json:"timestamp"`
	Quantity    int32           `db:"quantity" json:"quantity"`
	UnitPrice   decimal.Decimal `db:"unit_price" json:"unitPrice"`
	TotalAmount decimal.Decimal `db:"total_amount" json:"totalAmount"`
}
