package service

import (
	"projects-monitor/src/entity"
)

type ProjectsIndex struct {
	repository Repository
}

func newProjectsIndex(r Repository) *ProjectsIndex {
	return &ProjectsIndex{repository: r}
}

func (s *ProjectsIndex) GetAllProjects() (*entity.ProjectsList, error) {
	return s.repository.GetAll()
}
