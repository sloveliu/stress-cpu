package main

import (
	"log"
	"os"
	"stress-cpu/config"
	"stress-cpu/handlers"
	"stress-cpu/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.ReleaseMode
	}
	gin.SetMode(ginMode)

	router := gin.New()
	stressHandler := handlers.NewStressHandler()
	routes.Setup(router, cfg.APIKey, stressHandler)

	serverAddr := ":" + cfg.Port
	log.Printf("server startup port: %s\n", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("server startup failed: %v", err)
	}
}
