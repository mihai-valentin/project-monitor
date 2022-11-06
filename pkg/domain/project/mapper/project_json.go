package mapper

import "github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"

type ProjectJson struct {
}

func NewProjectJsonMapper() *ProjectJson {
	return &ProjectJson{}
}

func (m *ProjectJson) MapProjectToJson(p *entity.Project) map[string]any {
	return map[string]any{
		"id":          p.Id,
		"name":        p.Name,
		"description": p.Description,
	}
}
