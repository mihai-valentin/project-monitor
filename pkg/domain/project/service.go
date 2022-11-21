package project

import (
	"errors"
	com "github.com/mihai-valentin/projects-monitor/pkg/com/project"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"
)

type Repository interface {
	GetAll() *entity.ProjectsList
	GetById(id int) (*entity.Project, bool)
	Save(p *entity.Project) (*entity.Project, error)
	Update(id int, p *entity.Project) error
	DeleteById(id int)
}

type Project struct {
	mapper     *Mapper
	repository Repository
}

func NewService(m *Mapper, r Repository) *Project {
	return &Project{
		mapper:     m,
		repository: r,
	}
}

func (s *Project) GetAllProjects() *com.ProjectsList {
	pl := s.repository.GetAll()

	return s.mapper.MapProjectsListToData(pl)
}

func (s *Project) GetProjectById(id int) (*com.Project, error) {
	p, ok := s.repository.GetById(id)

	if !ok {
		return nil, errors.New("project not found")
	}

	return s.mapper.MapProjectToData(p), nil
}

func (s *Project) SaveProject(pd *com.Project) error {
	p := s.mapper.MapDataToProject(pd)
	if _, err := s.repository.Save(p); err != nil {
		return err
	}

	return nil
}

func (s *Project) UpdateProjectById(id int, pd *com.Project) error {
	if _, ok := s.repository.GetById(id); !ok {
		return errors.New("project not found")
	}

	p := s.mapper.MapDataToProject(pd)

	return s.repository.Update(id, p)
}

func (s *Project) DeleteProjectById(id int) error {
	if _, ok := s.repository.GetById(id); !ok {
		return errors.New("project not found")
	}

	s.repository.DeleteById(id)

	return nil
}
