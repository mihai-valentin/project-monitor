package repository

import "github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"

type InMemoryProject struct {
	projectsList *entity.ProjectsList
}

func New() *InMemoryProject {
	return &InMemoryProject{
		projectsList: entity.FakeProjectsList(10),
	}
}

func (r *InMemoryProject) GetAll() *entity.ProjectsList {
	return r.projectsList
}

func (r *InMemoryProject) GetById(id int) (*entity.Project, bool) {
	for _, project := range r.projectsList.All() {
		if project.Id == id {
			return project, true
		}
	}

	return nil, false
}

func (r *InMemoryProject) Save(p *entity.Project) (*entity.Project, error) {
	p.Id = r.projectsList.GetIndex() + 1
	r.projectsList.Add(p)

	return p, nil
}

func (r *InMemoryProject) Update(id int, p *entity.Project) error {
	project, ok := r.GetById(id)

	if !ok {
		return nil
	}

	project.Name = p.Name
	project.Description = p.Description

	return nil
}

func (r *InMemoryProject) DeleteById(id int) {
	project, ok := r.GetById(id)

	if !ok {
		return
	}

	r.projectsList.Remove(project)
}
