package repository

import "github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"

type MockInMemory struct {
	*InMemory
}

func Mock() *MockInMemory {
	return &MockInMemory{
		&InMemory{},
	}
}

func (m *MockInMemory) SetData(pl *entity.ProjectsList) {
	m.projectsList = pl
}
