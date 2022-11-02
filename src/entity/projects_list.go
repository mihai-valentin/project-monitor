package entity

type ProjectsList struct {
	projects map[int]*Project
}

func NewProjectsList() *ProjectsList {
	return &ProjectsList{
		projects: map[int]*Project{},
	}
}

func (pl *ProjectsList) All() map[int]*Project {
	return pl.projects
}

func (pl *ProjectsList) Add(p *Project) {
	pl.projects[p.Id] = p
}
