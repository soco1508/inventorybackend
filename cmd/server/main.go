package main

import (
	"backend/config"
	"backend/internal/api/routes"
	"backend/pkg/db"
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	config, err := config.NewParsedConfig()
	if err != nil {
		log.Fatalf("%v", err)
	}

	dbConfig := db.DBConfig{
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		Username: config.Database.Username,
		Password: config.Database.Password,
		Name:     config.Database.Name,
	}

	sqlxDb, err := db.SqlxInitDB(dbConfig)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer func() {
		if err := sqlxDb.Close(); err != nil {
			log.Fatalf("Error when closing db connection: %v", err)
		}
	}()

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowOrigins = []string{"https://dongtech.org"}
	corsCfg.AllowCredentials = true
	corsCfg.AllowHeaders = []string{"*"}

	router.Use(cors.New(corsCfg))
	routes.RegisterDashboard(router, sqlxDb)
	routes.RegisterProduct(router, sqlxDb)
	routes.RegisterExpense(router, sqlxDb)
	routes.RegisterUser(router, sqlxDb)

	port := os.Getenv("PORT")
	if port == "" {
		port = config.ServerPort
	}

	host := config.ServerHost
	if host == "" {
		host = "0.0.0.0"
	}

	server := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: router,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Server is running on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start the server: %v", err)
		}
	}()

	<-sigChan
	log.Println("Server is closing...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = server.Shutdown(ctx); err != nil {
		log.Printf("Failed to shutdown server: %v", err)
	} else {
		log.Print("Gracefully shutdown server")
	}
}
