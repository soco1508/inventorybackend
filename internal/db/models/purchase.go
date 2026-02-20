package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Purchase struct {
	PurchaseID string          `db:"purchase_id" json:"purchaseId"`
	ProductID  string          `db:"product_id" json:"productId"`
	Timestamp  time.Time       `db:"timestamp" json:"timestamp"`
	Quantity   int32           `db:"quantity" json:"quantity"`
	UnitCost   decimal.Decimal `db:"unit_cost" json:"unitCost"`
	TotalCost  decimal.Decimal `db:"total_cost" json:"totalCost"`
}
