package routes

import (
	"backend/internal/api/handler"
	"backend/internal/db/repository"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterDashboard(router *gin.Engine, db *sqlx.DB) {
	//repository
	productRepo := repository.NewProductRepository(db)
	saleSummaryRepo := repository.NewSaleSummaryRepository(db)
	purchaseSummaryRepo := repository.NewPurchaseSummaryRepository(db)
	expenseSummaryRepo := repository.NewExpenseSummaryRepository(db)
	expenseByCategoryRepo := repository.NewExpenseByCategoryRepository(db)

	//service
	productService := service.NewProductService(productRepo)
	saleSummaryService := service.NewSaleSummaryService(saleSummaryRepo)
	purchaseSummaryService := service.NewPurchaseSummaryService(purchaseSummaryRepo)
	expenseSummaryService := service.NewExpenseSummaryService(expenseSummaryRepo)
	expenseByCategoryService := service.NewExpenseByCategoryService(expenseByCategoryRepo)

	//handler
	dashboardHandler := handler.NewDashboardHandler(productService, saleSummaryService, purchaseSummaryService, expenseSummaryService, expenseByCategoryService)

	//routes
	router.GET("/dashboard", dashboardHandler.GetDashboardMetrics)
}
