package project

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"
	"strconv"
)

type Mapper interface {
	MapFormToProject(*entity.ProjectForm) *entity.Project
	MapProjectsListToJson(pl *entity.ProjectsList) []map[string]any
	MapProjectToJson(p *entity.Project) map[string]any
}

type Service interface {
	SaveProject(project *entity.Project) error
	DeleteProjectById(id int) error
	GetAllProjects() *entity.ProjectsList
	GetAllProjectsPaginated(page int, count int) *entity.ProjectsPaginatedList
	GetProjectById(id int) (*entity.Project, error)
	UpdateProjectById(id int, project *entity.Project) error
}

type Project struct {
	mapper  Mapper
	service Service
}

func New(m Mapper, s Service) *Project {
	return &Project{
		mapper:  m,
		service: s,
	}
}

func (c *Project) getIdParamFromContext(ctx *gin.Context) (int, error) {
	idRaw := ctx.Param("id")
	id, err := strconv.Atoi(idRaw)

	if err != nil {
		return 0, errors.New("path parameter id must be an int")
	}

	return id, nil
}

func (c *Project) bindProjectFormAndMapToProject(f *entity.ProjectForm, ctx *gin.Context) (*entity.Project, error) {
	if err := ctx.ShouldBind(f); err != nil {
		return nil, err
	}

	return c.mapper.MapFormToProject(f), nil
}
