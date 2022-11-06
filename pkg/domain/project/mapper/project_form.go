package mapper

import "github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"

type ProjectForm struct {
}

func NewProjectFormMapper() *ProjectForm {
	return &ProjectForm{}
}

func (m *ProjectForm) MapFormToProject(f *entity.ProjectForm) *entity.Project {
	return &entity.Project{
		Name:        f.Name,
		Description: f.Description,
	}
}
