package entity

import "fmt"

type ProjectsList struct {
	index    int
	projects []*Project
}

func NewProjectsList() *ProjectsList {
	return &ProjectsList{
		index:    0,
		projects: make([]*Project, 0),
	}
}

func FakeProjectsList(len int) *ProjectsList {
	projectsList := make([]*Project, 0)

	for i := 0; i < len; i++ {
		fakeProjectName := fmt.Sprintf("%d_project", i)
		fakeProject := NewProject(i+1, fakeProjectName, "Description...")

		projectsList = append(projectsList, fakeProject)
	}

	return &ProjectsList{index: len, projects: projectsList}
}

func (pl *ProjectsList) CloneEmpty() *ProjectsList {
	return &ProjectsList{
		index:    pl.index,
		projects: make([]*Project, 0),
	}
}

func (pl *ProjectsList) Count() int {
	return len(pl.projects)
}

func (pl *ProjectsList) GetIndex() int {
	return pl.index
}

func (pl *ProjectsList) All() []*Project {
	return pl.projects
}

func (pl *ProjectsList) Add(p Project) {
	pl.index += 1
	pl.projects = append(pl.projects, &p)
}

func (pl *ProjectsList) Remove(p *Project) {
	newProjectsList := make([]*Project, 0)

	for _, project := range pl.projects {
		if project.Id == p.Id {
			continue
		}

		newProjectsList = append(newProjectsList, project)
	}

	pl.projects = newProjectsList
}
