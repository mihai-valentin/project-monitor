package project

import (
	"errors"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"
)

type Repository interface {
	GetAll() *entity.ProjectsList
	GetById(id int) (*entity.Project, bool)
	Save(p entity.Project) (*entity.Project, error)
	Update(id int, p entity.Project) error
	DeleteById(id int)
	GetPaginated(page int, perPage int) *entity.ProjectsPaginatedList
}

type Project struct {
	repository Repository
}

func NewService(r Repository) *Project {
	return &Project{r}
}

func (s *Project) GetAllProjects() *entity.ProjectsList {
	return s.repository.GetAll()
}

func (s *Project) GetProjectById(id int) (*entity.Project, error) {
	p, ok := s.repository.GetById(id)

	if !ok {
		return nil, errors.New("project not found")
	}

	return p, nil
}

func (s *Project) SaveProject(p entity.Project) error {
	if _, err := s.repository.Save(p); err != nil {
		return err
	}

	return nil
}

func (s *Project) UpdateProjectById(id int, p entity.Project) error {
	if _, ok := s.repository.GetById(id); !ok {
		return errors.New("project not found")
	}

	return s.repository.Update(id, p)
}

func (s *Project) DeleteProjectById(id int) error {
	if _, ok := s.repository.GetById(id); !ok {
		return errors.New("project not found")
	}

	s.repository.DeleteById(id)

	return nil
}

func (s *Project) GetAllProjectsPaginated(page int, perPage int) *entity.ProjectsPaginatedList {
	return s.repository.GetPaginated(page, perPage)
}
