package service

import (
	"backend/internal/db/models"
	"backend/internal/db/repository"
	"context"
)

type SaleSummaryService interface {
	GetSaleSummary(ctx context.Context) ([]*models.SaleSummary, error)
}

type saleSummaryService struct {
	saleSummaryRepo repository.SaleSummaryRepository
}

func NewSaleSummaryService(saleSummaryRepo repository.SaleSummaryRepository) SaleSummaryService {
	return &saleSummaryService{saleSummaryRepo: saleSummaryRepo}
}

func (s *saleSummaryService) GetSaleSummary(ctx context.Context) ([]*models.SaleSummary, error) {
	return s.saleSummaryRepo.GetSaleSummary(ctx)
}
