package handler

import (
	"backend/internal/db/models"
	"backend/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type ExpenseByCategoryRes struct {
	models.ExpenseByCategory
	Amount decimal.Decimal
}

type ExpenseHandler struct {
	expenseByCategoryService service.ExpenseByCategoryService
}

func NewExpenseHandler(expenseByCategoryService service.ExpenseByCategoryService) *ExpenseHandler {
	return &ExpenseHandler{expenseByCategoryService: expenseByCategoryService}
}

func (h *ExpenseHandler) GetExpensesByCategory(c *gin.Context) {
	expensesByCategory, err := h.expenseByCategoryService.GetExpenseByCategory(c.Request.Context())
	if err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	expenseByCategoryRes := []ExpenseByCategoryRes{}
	for _, item := range expensesByCategory {
		expenseByCategoryRes = append(expenseByCategoryRes, ExpenseByCategoryRes{
			ExpenseByCategory: *item,
			Amount:            item.Amount,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"expenseByCategorySummary": expenseByCategoryRes,
	})
}
