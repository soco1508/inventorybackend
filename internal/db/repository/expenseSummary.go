package repository

import (
	"backend/internal/db/models"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ExpenseSummaryRepository interface {
	GetExpenseSummary(ctx context.Context) ([]*models.ExpenseSummary, error)
}

type expenseSummaryRepository struct {
	db *sqlx.DB
}

func NewExpenseSummaryRepository(db *sqlx.DB) ExpenseSummaryRepository {
	return &expenseSummaryRepository{db: db}
}

func (p *expenseSummaryRepository) GetExpenseSummary(ctx context.Context) ([]*models.ExpenseSummary, error) {
	sql := `SELECT * FROM expense_summary ORDER BY date DESC LIMIT 5`
	rows, err := p.db.QueryxContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("query ExpenseSummary err: %v", err)
	}
	defer rows.Close()

	expenseSummary := []*models.ExpenseSummary{}
	for rows.Next() {
		item := models.ExpenseSummary{}
		if err = rows.StructScan(&item); err != nil {
			return nil, fmt.Errorf("StructScan ExpenseSummary err: %v", err)
		}
		expenseSummary = append(expenseSummary, &item)
	}

	return expenseSummary, nil
}
