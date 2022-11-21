package project

import (
	com "github.com/mihai-valentin/projects-monitor/pkg/com/project"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"
)

type Mapper struct {
}

func NewMapper() *Mapper {
	return &Mapper{}
}

func (m *Mapper) MapDataToProject(pd *com.Project) *entity.Project {
	return entity.NewProject(pd.Id, pd.Name, pd.Description)
}

func (m *Mapper) MapProjectToData(p *entity.Project) *com.Project {
	return &com.Project{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
	}
}

func (m *Mapper) MapProjectsListToData(pl *entity.ProjectsList) *com.ProjectsList {
	pld := com.NewProjectsList()

	for _, p := range pl.All() {
		p := m.MapProjectToData(p)
		pld.Add(p)
	}

	return pld
}
