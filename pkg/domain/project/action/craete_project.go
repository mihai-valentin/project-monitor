package action

import (
	"github.com/gin-gonic/gin"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"
	"net/http"
)

type CreateProjectFormMapper interface {
	MapFormToProject(projectForm *entity.ProjectForm) *entity.Project
}

type StoreService interface {
	SaveProject(project *entity.Project) error
}

type CreateProject struct {
	mapper  CreateProjectFormMapper
	service StoreService
}

func NewCreateProjectAction(m CreateProjectFormMapper, s StoreService) *CreateProject {
	return &CreateProject{
		mapper:  m,
		service: s,
	}
}

func (a *CreateProject) LoadRoutes(gr *gin.Engine) {
	gr.POST("/projects", a.storeProject)
}

func (a *CreateProject) storeProject(ctx *gin.Context) {
	var projectForm entity.ProjectForm

	if err := ctx.ShouldBind(&projectForm); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	project := a.mapper.MapFormToProject(&projectForm)

	if err := a.service.SaveProject(project); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}
