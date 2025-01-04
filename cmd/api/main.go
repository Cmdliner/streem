package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

)

func main() {

	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("No .env file found")
	}

	gin.SetMode(os.Getenv("GIN_MODE"))

	// Use the default gin router
	r := gin.Default()


	// Server health checks
/* 	r.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": "The Hood is up Commandliner⚡"})
	})


	// Register endpoint
	r.POST("/register", func(ctx *gin.Context) {
		var newUser models.User
		if err := ctx.BindJSON(&newUser); err != nil {
			return
		}
		_, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		ctx.IndentedJSON(http.StatusCreated, gin.H{"new user": newUser})
	})

	// Login
	r.POST("/login", func (ctx *gin.Context)  {
		var user dtos.UserLogin
		if err := ctx.BindJSON(&user); err != nil {
			return
		}
		
		


	}) */

	// Run server
	r.Run(":8080")

}
