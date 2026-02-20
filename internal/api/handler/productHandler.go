package handler

import (
	"backend/internal/db/models"
	"backend/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type ProductRequest struct {
	ProductId     string          `json:"productId" binding:"required"`
	Name          string          `json:"name" binding:"required"`
	Price         decimal.Decimal `json:"price"`
	Rating        decimal.Decimal `json:"rating"`
	StockQuantity int32           `json:"stockQuantity" binding:"required"`
}

type ProductHandler struct {
	productSer service.ProductService
}

func NewProductHandler(productSer service.ProductService) *ProductHandler {
	return &ProductHandler{
		productSer: productSer,
	}
}

func (h *ProductHandler) FindMany(c *gin.Context) {
	search := c.Query("search")
	products, err := h.productSer.FindMany(c.Request.Context(), search)
	if err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	productReq := ProductRequest{}
	if err := c.ShouldBindJSON(&productReq); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	product := models.Product{
		ProductID:     productReq.ProductId,
		Name:          productReq.Name,
		Price:         productReq.Price,
		Rating:        productReq.Rating,
		StockQuantity: productReq.StockQuantity,
	}

	err := h.productSer.CreateProduct(c.Request.Context(), product)
	if err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Create product successfully!",
		"product": productReq,
	})
}
