package entity

type ProjectsPaginatedList struct {
	Items      []*Project
	Pagination *Pagination
}

func NewProjectsPaginatedList(i []*Project, p *Pagination) *ProjectsPaginatedList {
	return &ProjectsPaginatedList{
		Items:      i,
		Pagination: p,
	}
}
