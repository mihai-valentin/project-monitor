package provider

import (
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"
	"github.com/mihai-valentin/projects-monitor/pkg/transport/gin/controller/project"
)

type ProjectDomain interface {
	GetAllProjects() *entity.ProjectsList
	GetAllProjectsPaginated(page int, count int) *entity.ProjectsPaginatedList
	GetProjectById(id int) (*entity.Project, error)
	SaveProject(project *entity.Project) error
	UpdateProjectById(id int, project *entity.Project) error
	DeleteProjectById(id int) error
}

func (p *Gin) LoadProjectDomain(d ProjectDomain) {
	c := project.New(p.mapper, d)

	projects := p.router.Group("/projects")
	{
		projects.GET("", c.GetAllProjects)
		projects.GET("/pag", c.GetAllProjectsPaginated)
		projects.POST("", c.StoreProject)

		projects.GET("/:id", c.GetProjectById)
		projects.PUT("/:id", c.UpdateProject)
		projects.DELETE("/:id", c.DeleteProject)
	}
}
