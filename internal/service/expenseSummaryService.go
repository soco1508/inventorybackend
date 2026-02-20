package service

import (
	"backend/internal/db/models"
	"backend/internal/db/repository"
	"context"
)

type ExpenseSummaryService interface {
	GetExpenseSummary(ctx context.Context) ([]*models.ExpenseSummary, error)
}

type expenseSummaryService struct {
	expenseSummaryRepo repository.ExpenseSummaryRepository
}

func NewExpenseSummaryService(expenseSummaryRepo repository.ExpenseSummaryRepository) ExpenseSummaryService {
	return &expenseSummaryService{expenseSummaryRepo: expenseSummaryRepo}
}

func (s *expenseSummaryService) GetExpenseSummary(ctx context.Context) ([]*models.ExpenseSummary, error) {
	return s.expenseSummaryRepo.GetExpenseSummary(ctx)
}
