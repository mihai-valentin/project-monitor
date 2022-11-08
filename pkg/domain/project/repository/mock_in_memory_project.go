package repository

import "github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"

type MockInMemoryProject struct {
	*InMemoryProject
}

func Mock() *MockInMemoryProject {
	return &MockInMemoryProject{
		&InMemoryProject{},
	}
}

func (m *MockInMemoryProject) SetData(pl *entity.ProjectsList) {
	m.projectsList = pl
}
