package routes

import (
	"backend/internal/api/handler"
	"backend/internal/db/repository"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterProduct(router *gin.Engine, db *sqlx.DB) {
	//repository
	productRepo := repository.NewProductRepository(db)

	//service
	productSer := service.NewProductService(productRepo)

	//handler
	productHandler := handler.NewProductHandler(productSer)

	router.GET("/products", productHandler.FindMany)
	router.POST("/products", productHandler.CreateProduct)
}
