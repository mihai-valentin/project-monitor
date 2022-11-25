package project

import (
	"github.com/gin-gonic/gin"
	dto "github.com/mihai-valentin/projects-monitor/pkg/com/project"
	mapper "github.com/mihai-valentin/projects-monitor/pkg/transport/api/project"
	"net/http"
)

type projectsList struct {
	mapper  *mapper.Mapper
	service dto.Domain
}

func newProjectsListController(m *mapper.Mapper, s dto.Domain) *projectsList {
	return &projectsList{
		mapper:  m,
		service: s,
	}
}

func (c *projectsList) GetAllProjects(ctx *gin.Context) {
	allProjects := c.service.GetAllProjects()
	projects := c.mapper.MapProjectsListToJson(allProjects)

	ctx.JSON(http.StatusOK, projects)
}
