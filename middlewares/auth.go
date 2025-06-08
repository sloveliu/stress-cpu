package middlewares

import (
	"crypto/subtle"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse API 錯誤回應格式
type ErrorResponse struct {
	Error string `json:"error"`
}

func AuthMiddleware(expectedAPIKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")
		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "API key required"})
			c.Abort()
			return
		}

		if subtle.ConstantTimeCompare([]byte(apiKey), []byte(expectedAPIKey)) != 1 {
			c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Invalid API key"})
			c.Abort()
			return
		}

		c.Next()
	}
}
