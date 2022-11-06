package action

import (
	"github.com/gin-gonic/gin"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"
	"net/http"
)

type ProjectListJsonMapper interface {
	MapProjectsListToJson(pl *entity.ProjectsList) []map[string]any
}

type IndexService interface {
	GetAllProjects() *entity.ProjectsList
}

type GetAllProjects struct {
	mapper  ProjectListJsonMapper
	service IndexService
}

func NewGetAllProjectsAction(m ProjectListJsonMapper, s IndexService) *GetAllProjects {
	return &GetAllProjects{
		mapper:  m,
		service: s,
	}
}

func (a *GetAllProjects) LoadRoutes(gr *gin.Engine) {
	gr.GET("/projects", a.getAllProjects)
}

func (a *GetAllProjects) getAllProjects(ctx *gin.Context) {
	projects := a.mapper.MapProjectsListToJson(a.service.GetAllProjects())

	ctx.JSON(http.StatusOK, projects)
}
