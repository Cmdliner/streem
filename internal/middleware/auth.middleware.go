package middleware

import (
	"net/http"
	"strings"

	"github.com/Cmdliner/streem/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
			c.Abort()
			return
		}

		if len(authHeader) != 2 || strings.Split(authHeader, " ")[0] != "Bearer" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")[1]
		if bearerToken != "" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
		}

		token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
			if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(jwtSecret), nil
		})

		 if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		 }

		 if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["user_id"].(primitive.ObjectID)
			c.Set("user", &model.User{ID: userID})
			c.Next()
		 } else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		 }
	}
}