package service

import (
	"backend/internal/db/models"
	"backend/internal/db/repository"
	"context"
)

type ProductService interface {
	GetPopularProducts(ctx context.Context) ([]*models.Product, error)
	FindMany(ctx context.Context, search string) ([]*models.Product, error)
	CreateProduct(ctx context.Context, product models.Product) error
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

func (s *productService) GetPopularProducts(ctx context.Context) ([]*models.Product, error) {
	return s.productRepo.GetPopularProducts(ctx)
}

func (s *productService) FindMany(ctx context.Context, search string) ([]*models.Product, error) {
	return s.productRepo.FindMany(ctx, search)
}

func (s *productService) CreateProduct(ctx context.Context, product models.Product) error {
	return s.productRepo.CreateProduct(ctx, product)
}
