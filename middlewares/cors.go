package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
  config := cors.DefaultConfig()
  config.AllowOrigins = []string{"*"}
  config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
  config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
  return cors.New(config)
}
