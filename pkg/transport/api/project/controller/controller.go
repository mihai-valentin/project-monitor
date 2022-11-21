package project

import (
	"errors"
	"github.com/gin-gonic/gin"
	dto "github.com/mihai-valentin/projects-monitor/pkg/com/project"
	"github.com/mihai-valentin/projects-monitor/pkg/transport/api/project"
	"github.com/mihai-valentin/projects-monitor/pkg/transport/api/project/form"
	"strconv"
)

type Controller struct {
	mapper  *project.Mapper
	service dto.Domain
}

func NewController(m *project.Mapper, s dto.Domain) *Controller {
	return &Controller{
		mapper:  m,
		service: s,
	}
}

func (c *Controller) getIdParamFromContext(ctx *gin.Context) (int, error) {
	idRaw := ctx.Param("id")
	id, err := strconv.Atoi(idRaw)

	if err != nil {
		return 0, errors.New("path parameter id must be an int")
	}

	return id, nil
}

func (c *Controller) bindProjectFormAndMapToProject(f *form.Project, ctx *gin.Context) (*dto.Project, error) {
	if err := ctx.ShouldBind(f); err != nil {
		return nil, err
	}

	return c.mapper.MapFormToProject(f), nil
}
