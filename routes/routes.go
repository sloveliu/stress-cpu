package routes

import (
	"net/http"
	"stress-cpu/handlers"
	"stress-cpu/middlewares"

	"github.com/gin-gonic/gin"
)

func registerGetHead(router gin.IRoutes, path string, handler gin.HandlerFunc) {
	router.GET(path, handler)
	router.HEAD(path, handler)
}

func Setup(router *gin.Engine, apiKey string, stressHandler *handlers.StressHandler) {
	router.Use(middlewares.LoggerMiddleware())
	router.Use(gin.Recovery())

	registerGetHead(router, "/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "status ok"})
	})

	registerGetHead(router, "/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	authorized := router.Group("/stress")
	authorized.Use(middlewares.AuthMiddleware(apiKey))
	{
		authorized.GET("/start", stressHandler.Start)
		authorized.GET("/stop", stressHandler.Stop)
		authorized.GET("/status", stressHandler.Status)
	}
}
