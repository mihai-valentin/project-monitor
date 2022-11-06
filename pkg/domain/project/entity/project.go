package entity

type Project struct {
	Id          int
	Name        string
	Description string
}

func NewProject(id int, name string, description string) *Project {
	return &Project{
		Id:          id,
		Name:        name,
		Description: description,
	}
}
