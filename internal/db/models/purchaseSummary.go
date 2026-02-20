package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type PurchaseSummary struct {
	PurchaseSummaryID string           `db:"purchase_summary_id" json:"purchaseSummaryId"`
	TotalPurchased    decimal.Decimal  `db:"total_purchased" json:"totalPurchased"`
	ChangePercentage  *decimal.Decimal `db:"change_percentage" json:"changePercentage"`
	Date              time.Time        `db:"date" json:"date"`
}
