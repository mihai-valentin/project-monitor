package entity

type Project struct {
	Id   int
	Name string
}

func NewProject(id int, name string) *Project {
	return &Project{
		Id:   id,
		Name: name,
	}
}
