package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/mihai-valentin/projects-monitor/pkg/transport/rest/middleware"
	"github.com/mihai-valentin/projects-monitor/pkg/transport/rest/router"
)

func RegisterTransport() *gin.Engine {
	m := middleware.New()
	r := router.New(m)

	return r.InitGinRouter()
}
