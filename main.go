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
	// 信任這些 IP 代理來源，以確保 gin 不會直接將代理 IP 作為客戶端
	err := router.SetTrustedProxies([]string{"127.0.0.1", "::1", "172.17.0.0/16"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	stressHandler := handlers.NewStressHandler()
	routes.Setup(router, cfg.APIKey, stressHandler)

	serverAddr := ":" + cfg.Port
	log.Printf("server startup port: %s\n", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("server startup failed: %v", err)
	}
}
