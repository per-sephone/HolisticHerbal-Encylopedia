package logger

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func LogRequestContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		entry := log.WithFields(log.Fields{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		})

		// Store the entry in the context so you can reuse it later
		c.Set("logger", entry)
		c.Next()
	}
}
