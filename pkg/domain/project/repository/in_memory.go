package repository

import (
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"
	"math"
)

type InMemory struct {
	projectsList *entity.ProjectsList
}

func New() *InMemory {
	return &InMemory{
		projectsList: entity.NewFakeProjectsList(10),
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

func (r *InMemory) GetPaginated(page int, perPage int) *entity.ProjectsPaginatedList {
	startIndex := (page - 1) * perPage
	endIndex := startIndex + perPage

	totalPages := math.Ceil(float64(r.projectsList.Count()) / float64(perPage))

	ppl := entity.NewProjectsPaginatedList(perPage, int(totalPages), page)

	for i := startIndex; i < endIndex; i++ {
		p, ok := r.projectsList.Get(i)

		if !ok {
			continue
		}

		ppl.Items = append(ppl.Items, p)
	}

	return ppl
}
