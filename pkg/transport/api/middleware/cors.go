package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	allowOrigin      = "*"
	allowCredentials = "true"
	allowHeaders     = "Content-Type, Content-Length, Accept-Encoding, Accept, Origin, Cache-Control"
	allowMethods     = "POST, OPTIONS, GET, PUT, DELETE"
)

func (m *Middleware) Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", allowOrigin)
	c.Header("Access-Control-Allow-Credentials", allowCredentials)
	c.Header("Access-Control-Allow-Headers", allowHeaders)
	c.Header("Access-Control-Allow-Methods", allowMethods)

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)

		return
	}

	c.Next()
}
