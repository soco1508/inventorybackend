package models

import "github.com/shopspring/decimal"

type Product struct {
	ProductID     string          `db:"product_id" json:"productId"`
	Name          string          `db:"name" json:"name"`
	Price         decimal.Decimal `db:"price" json:"price"`
	Rating        decimal.Decimal `db:"rating" json:"rating"`
	StockQuantity int32           `db:"stock_quantity" json:"stockQuantity"`
}
