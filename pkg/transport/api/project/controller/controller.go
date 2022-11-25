package project

import (
	"errors"
	"github.com/gin-gonic/gin"
	dto "github.com/mihai-valentin/projects-monitor/pkg/com/project"
	mapper "github.com/mihai-valentin/projects-monitor/pkg/transport/api/project"
	"github.com/mihai-valentin/projects-monitor/pkg/transport/api/project/form"
	"strconv"
)

type Controller struct {
	*project
	*projectsList
}

func NewController(m *mapper.Mapper, s dto.Domain) *Controller {
	return &Controller{
		project:      newProjectController(m, s),
		projectsList: newProjectsListController(m, s),
	}
}

func getIdParamFromContext(ctx *gin.Context) (int, error) {
	idRaw := ctx.Param("id")
	id, err := strconv.Atoi(idRaw)

	if err != nil {
		return 0, errors.New("path parameter id must be an int")
	}

	return id, nil
}

func bindProjectFormAndMapToProject(ctx *gin.Context, f *form.Project, m *mapper.Mapper) (*dto.Project, error) {
	if err := ctx.ShouldBind(f); err != nil {
		return nil, err
	}

	return m.MapFormToProject(f), nil
}
