package project

import (
	"github.com/gin-gonic/gin"
	com "github.com/mihai-valentin/projects-monitor/pkg/com/project"
	"github.com/mihai-valentin/projects-monitor/pkg/transport/api/project/controller"
)

type Provider struct {
	controller *project.Controller
}

func New(s com.Domain) *Provider {
	m := NewProjectMapper()
	c := project.NewController(m, s)

	return &Provider{
		controller: c,
	}
}

func (p *Provider) RegisterRoutes(gr *gin.Engine) {
	projects := gr.Group("/projects")
	{
		projects.GET("", p.controller.GetAllProjects)
		projects.POST("", p.controller.StoreProject)

		projects.GET("/:id", p.controller.GetProjectById)
		projects.PUT("/:id", p.controller.UpdateProject)
		projects.DELETE("/:id", p.controller.DeleteProject)
	}
}
