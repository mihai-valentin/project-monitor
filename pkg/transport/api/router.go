package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mihai-valentin/projects-monitor/pkg/transport/api/middleware"
)

type Controller interface {
	RegisterRoutes(engine *gin.Engine)
}

type Router struct {
	engine *gin.Engine
}

func New() *Router {
	m := middleware.New()

	gr := gin.Default()
	gr.Use(m.Cors)

	return &Router{engine: gr}
}

func (r *Router) Run() error {
	return r.engine.Run()
}

func (r *Router) RegisterController(c Controller) {
	c.RegisterRoutes(r.engine)
}
