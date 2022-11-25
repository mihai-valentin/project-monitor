package entity

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

func (pl *ProjectsList) Get(i int) (*Project, bool) {
	if len(pl.projects) < i {
		return nil, false
	}

	return pl.projects[i], true
}
