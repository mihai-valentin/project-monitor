package entity

type ProjectsPaginatedList struct {
	Items       []*Project `json:"items"`
	PerPage     int        `json:"per_page"`
	TotalPages  int        `json:"total_pages"`
	CurrentPage int        `json:"current_page"`
}

func NewProjectsPaginatedList(pp int, tp int, cp int) *ProjectsPaginatedList {
	return &ProjectsPaginatedList{
		Items:       make([]*Project, 0),
		PerPage:     pp,
		TotalPages:  tp,
		CurrentPage: cp,
	}
}
