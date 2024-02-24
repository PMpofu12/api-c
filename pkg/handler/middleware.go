package handler

import (
	"github.com/gin-gonic/gin"
)

func LoggingMiddleware(c *gin.Context) {
	println("Request received:", c.Request.Method, c.Request.URL)
	c.Next()
}
