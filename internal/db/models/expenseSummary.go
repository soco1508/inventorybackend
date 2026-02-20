package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type ExpenseSummary struct {
	ExpenseSummaryId string          `db:"expense_summary_id" json:"expenseSummaryId"`
	TotalExpenses    decimal.Decimal `db:"total_expenses" json:"totalExpenses"`
	Date             time.Time       `db:"date" json:"date"`
}
