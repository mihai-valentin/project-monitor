package action

import (
	"github.com/gin-gonic/gin"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"
	"net/http"
	"strconv"
)

type ProjectJsonMapper interface {
	MapProjectToJson(p *entity.Project) map[string]any
}

type ShowService interface {
	GetProjectById(id int) (*entity.Project, error)
}

type GetProject struct {
	mapper  ProjectJsonMapper
	service ShowService
}

func NewGetProjectAction(m ProjectJsonMapper, s ShowService) *GetProject {
	return &GetProject{m, s}
}

func (a *GetProject) LoadRoutes(gr *gin.Engine) {
	gr.GET("/projects/:id", a.getProjectById)
}

func (a *GetProject) getProjectById(ctx *gin.Context) {
	idRaw := ctx.Param("id")
	id, err := strconv.Atoi(idRaw)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": "path parameter id must be an int",
		})

		return
	}

	project, err := a.service.GetProjectById(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "project not found",
		})

		return
	}

	ctx.JSON(http.StatusOK, a.mapper.MapProjectToJson(project))
}
