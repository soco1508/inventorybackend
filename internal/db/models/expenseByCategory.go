package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type ExpenseByCategory struct {
	ExpenseByCategoryId string          `db:"expense_by_category_id" json:"expenseByCategoryId"`
	ExpenseSummaryId    string          `db:"expense_summary_id" json:"expenseSummaryId"`
	Date                time.Time       `db:"date" json:"date"`
	Category            string          `db:"category" json:"category"`
	Amount              decimal.Decimal `db:"amount" json:"amount"`
}
