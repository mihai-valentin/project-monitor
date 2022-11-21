package project

type ProjectsList struct {
	items []*Project
}

func NewProjectsList() *ProjectsList {
	return &ProjectsList{
		items: make([]*Project, 0),
	}
}

func (pl *ProjectsList) Count() int {
	return len(pl.items)
}

func (pl *ProjectsList) All() []*Project {
	return pl.items
}

func (pl *ProjectsList) Add(p *Project) {
	pl.items = append(pl.items, p)
}
