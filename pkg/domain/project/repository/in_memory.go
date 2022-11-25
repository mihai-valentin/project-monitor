package repository

import "github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"

type InMemory struct {
	projectsList *entity.ProjectsList
}

func New() *InMemory {
	return &InMemory{
		projectsList: entity.FakeProjectsList(10),
	}
}

func (r *InMemory) GetAll() *entity.ProjectsList {
	return r.projectsList
}

func (r *InMemory) GetById(id int) (*entity.Project, bool) {
	for _, project := range r.projectsList.All() {
		if project.Id == id {
			return project, true
		}
	}

	return nil, false
}

func (r *InMemory) Save(p entity.Project) (*entity.Project, error) {
	p.Id = r.projectsList.GetIndex() + 1
	r.projectsList.Add(p)

	return &p, nil
}

func (r *InMemory) Update(id int, p entity.Project) error {
	project, ok := r.GetById(id)

	if !ok {
		return nil
	}

	project.Name = p.Name
	project.Description = p.Description

	return nil
}

func (r *InMemory) DeleteById(id int) {
	project, ok := r.GetById(id)

	if !ok {
		return
	}

	r.projectsList.Remove(project)
}
