package project

type Domain interface {
	SaveProject(project *Project) error
	DeleteProjectById(id int) error
	GetAllProjects() *ProjectsList
	GetProjectById(id int) (*Project, error)
	UpdateProjectById(id int, project *Project) error
}
