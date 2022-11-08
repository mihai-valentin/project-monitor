package entity

type Pagination struct {
	TotalPages  int
	CurrentPage int
	PerPage     int
}

func NewPagination(tp int, cp int, pp int) *Pagination {
	return &Pagination{
		TotalPages:  tp,
		CurrentPage: cp,
		PerPage:     pp,
	}
}
