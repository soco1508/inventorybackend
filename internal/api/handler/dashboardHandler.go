package handler

import (
	"backend/internal/service"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	productService           service.ProductService
	saleSummaryService       service.SaleSummaryService
	purchaseSummaryService   service.PurchaseSummaryService
	expenseSummaryService    service.ExpenseSummaryService
	expenseByCategoryService service.ExpenseByCategoryService
}

func NewDashboardHandler(
	productService service.ProductService,
	saleSummaryService service.SaleSummaryService,
	purchaseSummaryService service.PurchaseSummaryService,
	expenseSummaryService service.ExpenseSummaryService,
	expenseByCategoryService service.ExpenseByCategoryService,
) *DashboardHandler {
	return &DashboardHandler{
		productService:           productService,
		saleSummaryService:       saleSummaryService,
		purchaseSummaryService:   purchaseSummaryService,
		expenseSummaryService:    expenseSummaryService,
		expenseByCategoryService: expenseByCategoryService,
	}
}

func (h *DashboardHandler) GetDashboardMetrics(c *gin.Context) {
	ctx := c.Request.Context()

	start := time.Now()

	popularProducts, err := h.productService.GetPopularProducts(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Printf("%v", err)
		return
	}

	saleSummary, err := h.saleSummaryService.GetSaleSummary(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Printf("%v", err)
		return
	}

	purchaseSummary, err := h.purchaseSummaryService.GetPurchaseSummary(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Printf("%v", err)
		return
	}

	expenseSummary, err := h.expenseSummaryService.GetExpenseSummary(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Printf("%v", err)
		return
	}

	expenseByCategory, err := h.expenseByCategoryService.GetExpenseByCategory(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Printf("%v", err)
		return
	}

	elapsed := time.Since(start)
	log.Printf("Elapsed time: %v", elapsed)
	c.JSON(http.StatusOK, gin.H{
		"popularProducts":   popularProducts,
		"salesSummary":      saleSummary,
		"purchaseSummary":   purchaseSummary,
		"expenseSummary":    expenseSummary,
		"expenseByCategory": expenseByCategory,
	})
}
