package routes

import (
	"backend/internal/api/handler"
	"backend/internal/db/repository"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterUser(router *gin.Engine, db *sqlx.DB) {
	//repository
	userRepo := repository.NewUserRepository(db)

	//service
	userSer := service.NewUserService(userRepo)

	//handler
	userHandler := handler.NewUserHandler(userSer)

	router.GET("/users", userHandler.GetUsers)
	router.GET("/users/email", userHandler.GetUserByEmail)
}
