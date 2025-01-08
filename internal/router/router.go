package router

import (
	"net/http"

	"github.com/Cmdliner/streem/internal/handler"
	"github.com/gin-gonic/gin"
)


func SetupRouter(authHandler *handler.AuthHandler) *gin.Engine {
	r := gin.Default()

	// Register healthcheck handler
	r.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": "The Hood is up Commandliner⚡"})
	})

	// Group /api/v1 routes
	v1 := r.Group("/api/v1")
	
	// Mount auth endpoints
	auth := v1.Group("/auth") 
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	return r
}