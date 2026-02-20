package repository

import (
	"backend/internal/db/models"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ExpenseByCategoryRepository interface {
	GetExpenseByCategory(ctx context.Context) ([]*models.ExpenseByCategory, error)
}

type expenseByCategoryRepository struct {
	db *sqlx.DB
}

func NewExpenseByCategoryRepository(db *sqlx.DB) ExpenseByCategoryRepository {
	return &expenseByCategoryRepository{db: db}
}

func (r *expenseByCategoryRepository) GetExpenseByCategory(ctx context.Context) ([]*models.ExpenseByCategory, error) {
	sql := `SELECT expense_by_category_id,
				   expense_summary_id,
				   date,
				   category,
				   amount		
			FROM expense_by_category 
			ORDER BY date DESC 
			LIMIT 5
		`

	expenseByCategory := []*models.ExpenseByCategory{}
	if err := r.db.SelectContext(ctx, &expenseByCategory, sql); err != nil {
		return nil, fmt.Errorf("query ExpenseByCategory err: %v", err)
	}

	return expenseByCategory, nil
}
