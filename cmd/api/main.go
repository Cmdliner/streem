package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Cmdliner/streem/internal/config"
	"github.com/Cmdliner/streem/internal/database"
	"github.com/Cmdliner/streem/internal/handler"
	"github.com/Cmdliner/streem/internal/repository"
	"github.com/Cmdliner/streem/internal/router"
	"github.com/Cmdliner/streem/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))

	// Load App-wide config
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	client, err := database.Connect(cfg.MongoDB.URI)

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	defer client.Disconnect(context.Background())

	// Initialize repositories
	userRepo := repository.NewUserRepository(client.Database(cfg.MongoDB.Name))

	// Initialize services
	authService := service.NewAuthService(userRepo, cfg.JWT.Secret, time.Duration(cfg.JWT.ExpirationHours))

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService)

	// Setup router
	r := router.SetupRouter(authHandler)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(cfg.Server.Port),
		Handler: r,
	}

	// Start server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
    		// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
