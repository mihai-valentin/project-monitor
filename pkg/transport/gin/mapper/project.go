package mapper

import "github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"

type Project struct {
}

func newProjectMapper() *Project {
	return &Project{}
}

func (m *Project) MapFormToProject(f *entity.ProjectForm) *entity.Project {
	return &entity.Project{
		Name:        f.Name,
		Description: f.Description,
	}
}

func (m *Project) MapProjectToJson(p *entity.Project) map[string]any {
	return map[string]any{
		"id":          p.Id,
		"name":        p.Name,
		"description": p.Description,
	}
}

func (m *Project) MapProjectsListToJson(pl *entity.ProjectsList) []map[string]any {
	projectsListJson := make([]map[string]any, pl.Count())

	for i, project := range pl.All() {
		projectsListJson[i] = m.MapProjectToJson(project)
	}

	return projectsListJson
}
