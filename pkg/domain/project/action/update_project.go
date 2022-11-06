package action

import (
	"github.com/gin-gonic/gin"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"
	"net/http"
	"strconv"
)

type UpdateProjectFormMapper interface {
	MapFormToProject(projectForm *entity.ProjectForm) *entity.Project
}

type UpdateService interface {
	UpdateProjectById(id int, project *entity.Project) error
}

type UpdateProject struct {
	mapper  UpdateProjectFormMapper
	service UpdateService
}

func NewUpdateProjectAction(m UpdateProjectFormMapper, s UpdateService) *UpdateProject {
	return &UpdateProject{
		mapper:  m,
		service: s,
	}
}

func (a *UpdateProject) LoadRoutes(gr *gin.Engine) {
	gr.PUT("/projects/:id", a.updateProject)
}

func (a *UpdateProject) updateProject(ctx *gin.Context) {
	idRaw := ctx.Param("id")
	id, err := strconv.Atoi(idRaw)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": "path parameter id must be an int",
		})

		return
	}

	var projectForm entity.ProjectForm

	if err := ctx.ShouldBind(&projectForm); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	project := a.mapper.MapFormToProject(&projectForm)

	if err := a.service.UpdateProjectById(id, project); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}
