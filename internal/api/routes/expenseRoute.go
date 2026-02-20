package routes

import (
	"backend/internal/api/handler"
	"backend/internal/db/repository"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterExpense(router *gin.Engine, db *sqlx.DB) {
	//repository
	expenseRepo := repository.NewExpenseByCategoryRepository(db)

	//service
	expenseSer := service.NewExpenseByCategoryService(expenseRepo)

	//handler
	expenseHandler := handler.NewExpenseHandler(expenseSer)

	router.GET("/expenses", expenseHandler.GetExpensesByCategory)
}
