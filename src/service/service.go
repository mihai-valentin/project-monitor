package service

import (
	"projects-monitor/src/entity"
)

type Repository interface {
	GetAll() (*entity.ProjectsList, error)
}

type Service struct {
	*ProjectsIndex
}

func New(r Repository) *Service {
	return &Service{
		ProjectsIndex: newProjectsIndex(r),
	}
}
