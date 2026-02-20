package service

import (
	"backend/internal/db/models"
	"backend/internal/db/repository"
	"context"
)

type ExpenseByCategoryService interface {
	GetExpenseByCategory(ctx context.Context) ([]*models.ExpenseByCategory, error)
}

type expenseByCategoryService struct {
	expenseByCategoryRepo repository.ExpenseByCategoryRepository
}

func NewExpenseByCategoryService(expenseByCategoryRepo repository.ExpenseByCategoryRepository) ExpenseByCategoryService {
	return &expenseByCategoryService{expenseByCategoryRepo: expenseByCategoryRepo}
}

func (s *expenseByCategoryService) GetExpenseByCategory(ctx context.Context) ([]*models.ExpenseByCategory, error) {
	return s.expenseByCategoryRepo.GetExpenseByCategory(ctx)
}
