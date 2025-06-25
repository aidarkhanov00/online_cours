package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"project-online-course/middlewares"
)

// JWTMiddleware проверяет токен в каждом запросе
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен не найден"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := middlewares.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			c.Abort()
			return
		}

		// Добавляем данные из токена в контекст
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}
