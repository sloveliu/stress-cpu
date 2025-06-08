package middlewares

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {
			var timestampFormatted string
			tzEnv := os.Getenv("TZ")
			loc, err := time.LoadLocation(tzEnv)
			if err != nil {
				log.Printf("failed to load location for timezone '%s': %v. use default setting", tzEnv, err)
				timestampFormatted = params.TimeStamp.Format(time.DateTime)
			} else {
				timestampFormatted = params.TimeStamp.In(loc).Format(time.DateTime)
			}

			logEntry := gin.H{
				"client_ip":  params.ClientIP,
				"timestamp":  timestampFormatted,
				"method":     params.Method,
				"path":       params.Path,
				"status":     params.StatusCode,
				"latency_ms": params.Latency.Milliseconds(),
			}

			if params.ErrorMessage != "" {
				logEntry["error"] = params.ErrorMessage
			}

			logBytes, _ := json.Marshal(logEntry)
			return string(logBytes) + "\n"
		},
	})
}
