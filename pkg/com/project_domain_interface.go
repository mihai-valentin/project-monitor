package com

import "github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"

type Project interface {
	SaveProject(p entity.Project) error
	DeleteProjectById(id int) error
	GetAllProjects() *entity.ProjectsList
	GetProjectById(id int) (*entity.Project, error)
	UpdateProjectById(id int, p entity.Project) error
	GetAllProjectsPaginated(page int, perPage int) *entity.ProjectsPaginatedList
}
