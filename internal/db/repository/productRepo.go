package repository

import (
	"backend/internal/db/models"
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ProductRepository interface {
	GetPopularProducts(ctx context.Context) ([]*models.Product, error)
	FindMany(ctx context.Context, search string) ([]*models.Product, error)
	CreateProduct(ctx context.Context, product models.Product) error
}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepository{db: db}
}

func (p *productRepository) GetPopularProducts(ctx context.Context) ([]*models.Product, error) {
	sql := `SELECT * FROM products ORDER BY stock_quantity DESC LIMIT 15`
	rows, err := p.db.QueryxContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("query products err:\n %+v", err)
	}
	defer rows.Close()

	products := []*models.Product{}
	for rows.Next() {
		product := models.Product{}
		if err = rows.StructScan(&product); err != nil {
			return nil, fmt.Errorf("StructScan product err:\n %+v", err)
		}
		products = append(products, &product)
	}

	return products, nil
}

func (p *productRepository) FindMany(ctx context.Context, search string) ([]*models.Product, error) {
	query := `SELECT * FROM products WHERE LOWER(name) LIKE $1`
	parsedSearch := strings.ToLower(strings.TrimSpace(search))
	rows, err := p.db.QueryxContext(ctx, query, "%"+parsedSearch+"%")
	if err != nil {
		return nil, fmt.Errorf("find products err:\n %+v", err)
	}
	defer rows.Close()

	products := []*models.Product{}
	for rows.Next() {
		product := models.Product{}
		if err := rows.StructScan(&product); err != nil {
			return nil, fmt.Errorf("StructScan product err:\n %+v", err)
		}
		products = append(products, &product)
	}

	return products, nil
}

func (p *productRepository) CreateProduct(ctx context.Context, product models.Product) error {
	sql := `INSERT INTO products(product_id, name, price, rating, stock_quantity)
			VALUES (:product_id, :name, :price, :rating, :stock_quantity)`

	_, err := p.db.NamedExecContext(ctx, sql, &product)
	if err != nil {
		return fmt.Errorf("failed to create product:\n %+v", err)
	}

	return nil
}
