package entity

type Project struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewProject(id int, name string, description string) *Project {
	return &Project{
		Id:          id,
		Name:        name,
		Description: description,
	}
}
