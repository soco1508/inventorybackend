package handler

import (
	"backend/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetUsers(c.Request.Context())
	if err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	var userReq UserReq
	if err := c.BindJSON(&userReq); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	rs, err := h.userService.GetUserByEmail(c.Request.Context(), userReq.Email, userReq.Name)
	if err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": rs})
}

type UserReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
