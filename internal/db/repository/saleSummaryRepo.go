package repository

import (
	"backend/internal/db/models"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SaleSummaryRepository interface {
	GetSaleSummary(ctx context.Context) ([]*models.SaleSummary, error)
}

type saleSummaryRepository struct {
	db *sqlx.DB
}

func NewSaleSummaryRepository(db *sqlx.DB) SaleSummaryRepository {
	return &saleSummaryRepository{db: db}
}

func (p *saleSummaryRepository) GetSaleSummary(ctx context.Context) ([]*models.SaleSummary, error) {
	sql := `SELECT * FROM sales_summary ORDER BY date DESC LIMIT 5`
	rows, err := p.db.QueryxContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("query SalesSummary err: %v", err)
	}
	defer rows.Close()

	saleSummary := []*models.SaleSummary{}
	for rows.Next() {
		item := models.SaleSummary{}
		if err = rows.StructScan(&item); err != nil {
			return nil, fmt.Errorf("StructScan SalesSummary err: %v", err)
		}
		saleSummary = append(saleSummary, &item)
	}

	return saleSummary, nil
}
