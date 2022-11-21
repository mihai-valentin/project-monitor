package project

import (
	"github.com/mihai-valentin/projects-monitor/pkg/com/project"
	"github.com/mihai-valentin/projects-monitor/pkg/transport/api/project/form"
)

type Mapper struct {
}

func NewProjectMapper() *Mapper {
	return &Mapper{}
}

func (m *Mapper) MapFormToProject(f *form.Project) *project.Project {
	return &project.Project{
		Name:        f.Name,
		Description: f.Description,
	}
}

func (m *Mapper) MapProjectToJson(p *project.Project) map[string]any {
	return map[string]any{
		"id":          p.Id,
		"name":        p.Name,
		"description": p.Description,
	}
}

func (m *Mapper) MapProjectsListToJson(pl *project.ProjectsList) []map[string]any {
	projectsListJson := make([]map[string]any, pl.Count())

	for i, p := range pl.All() {
		projectsListJson[i] = m.MapProjectToJson(p)
	}

	return projectsListJson
}
