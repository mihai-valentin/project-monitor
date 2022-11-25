package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mihai-valentin/projects-monitor/pkg/com"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"
	"github.com/mihai-valentin/projects-monitor/pkg/transport/api/form"
	"net/http"
)

type ProjectsList struct {
	service com.Project
}

func NewProjectsListController(s com.Project) *ProjectsList {
	return &ProjectsList{s}
}

func (c *ProjectsList) RegisterRoutes(gr *gin.Engine) {
	projects := gr.Group("/projects")
	{
		projects.GET("", c.getAllProjects)
		projects.POST("", c.storeProject)
	}
}

func (c *ProjectsList) storeProject(ctx *gin.Context) {
	var pf *form.Project

	if err := ctx.ShouldBind(pf); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	p := entity.Project{
		Name:        pf.Name,
		Description: pf.Description,
	}

	if err := c.service.SaveProject(p); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (c *ProjectsList) getAllProjects(ctx *gin.Context) {
	pl := c.service.GetAllProjects()

	ctx.JSON(http.StatusOK, pl)
}
