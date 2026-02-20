package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type SaleSummary struct {
	SalesSummaryID   string           `db:"sales_summary_id" json:"salesSummaryId"`
	TotalValue       decimal.Decimal  `db:"total_value" json:"totalValue"`
	ChangePercentage *decimal.Decimal `db:"change_percentage,omitempty" json:"changePercentage,omitempty"`
	Date             time.Time        `db:"date" json:"date"`
}
