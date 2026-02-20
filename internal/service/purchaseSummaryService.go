package service

import (
	"backend/internal/db/models"
	"backend/internal/db/repository"
	"context"
)

type PurchaseSummaryService interface {
	GetPurchaseSummary(ctx context.Context) ([]*models.PurchaseSummary, error)
}

type purchaseSummaryService struct {
	purchaseSummaryRepo repository.PurchaseSummaryRepository
}

func NewPurchaseSummaryService(purchaseSummaryRepo repository.PurchaseSummaryRepository) PurchaseSummaryService {
	return &purchaseSummaryService{purchaseSummaryRepo: purchaseSummaryRepo}
}

func (s *purchaseSummaryService) GetPurchaseSummary(ctx context.Context) ([]*models.PurchaseSummary, error) {
	return s.purchaseSummaryRepo.GetPurchaseSummary(ctx)
}
