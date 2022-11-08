package service

import (
	"errors"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"
)

type Repository interface {
	GetAll() *entity.ProjectsList
	GetAllPaginated(page int, count int) *entity.ProjectsPaginatedList
	GetById(id int) (*entity.Project, bool)
	Save(p *entity.Project) (*entity.Project, error)
	Update(id int, p *entity.Project) error
	DeleteById(id int)
}

type Project struct {
	repository Repository
}

func New(r Repository) *Project {
	return &Project{r}
}

func (s *Project) GetAllProjects() *entity.ProjectsList {
	return s.repository.GetAll()
}

func (s *Project) GetAllProjectsPaginated(page int, count int) *entity.ProjectsPaginatedList {
	return s.repository.GetAllPaginated(page, count)
}

func (s *Project) GetProjectById(id int) (*entity.Project, error) {
	project, ok := s.repository.GetById(id)

	if !ok {
		return nil, errors.New("project not found")
	}

	return project, nil
}

func (s *Project) SaveProject(project *entity.Project) error {
	if _, err := s.repository.Save(project); err != nil {
		return err
	}

	return nil
}

func (s *Project) UpdateProjectById(id int, project *entity.Project) error {
	if _, ok := s.repository.GetById(id); !ok {
		return errors.New("project not found")
	}

	return s.repository.Update(id, project)
}

func (s *Project) DeleteProjectById(id int) error {
	if _, ok := s.repository.GetById(id); !ok {
		return errors.New("project not found")
	}

	s.repository.DeleteById(id)

	return nil
}
