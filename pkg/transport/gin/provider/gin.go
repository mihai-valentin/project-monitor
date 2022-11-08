package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/mihai-valentin/projects-monitor/pkg/transport/gin/mapper"
	"github.com/mihai-valentin/projects-monitor/pkg/transport/gin/middleware"
)

type Gin struct {
	router *gin.Engine
	mapper *mapper.Mapper
}

func New() *Gin {
	m := middleware.New()

	gr := gin.Default()
	gr.Use(m.CORS)

	return &Gin{
		router: gr,
		mapper: mapper.New(),
	}
}

func (p *Gin) Run() error {
	return p.router.Run()
}
