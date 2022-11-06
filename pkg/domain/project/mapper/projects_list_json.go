package mapper

import "github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"

type ProjectListJson struct {
	ProjectJson
}

func NewProjectListJsonMapper() *ProjectListJson {
	return &ProjectListJson{}
}

func (m *ProjectListJson) MapProjectsListToJson(pl *entity.ProjectsList) []map[string]any {
	projectsListJson := make([]map[string]any, pl.Count())

	for i, project := range pl.All() {
		projectsListJson[i] = m.MapProjectToJson(project)
	}

	return projectsListJson
}
