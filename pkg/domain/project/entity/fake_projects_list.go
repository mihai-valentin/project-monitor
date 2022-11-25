package entity

import "fmt"

type FakeProjectsList struct {
	*ProjectsList
}

func NewFakeProjectsList(len int) *ProjectsList {
	projectsList := make([]*Project, 0)

	for i := 0; i < len; i++ {
		fakeProjectName := fmt.Sprintf("%d_project", i)
		fakeProject := NewProject(i+1, fakeProjectName, "Description...")

		projectsList = append(projectsList, fakeProject)
	}

	return &ProjectsList{index: len, projects: projectsList}
}
