package project

import (
	"github.com/gin-gonic/gin"
	dto "github.com/mihai-valentin/projects-monitor/pkg/com/project"
	mapper "github.com/mihai-valentin/projects-monitor/pkg/transport/api/project"
	"github.com/mihai-valentin/projects-monitor/pkg/transport/api/project/form"
	"net/http"
)

type project struct {
	mapper  *mapper.Mapper
	service dto.Domain
}

func newProjectController(m *mapper.Mapper, s dto.Domain) *project {
	return &project{
		mapper:  m,
		service: s,
	}
}

func (c *project) GetProjectById(ctx *gin.Context) {
	id, err := getIdParamFromContext(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	p, err := c.service.GetProjectById(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, c.mapper.MapProjectToJson(p))
}

func (c *project) StoreProject(ctx *gin.Context) {
	var projectForm *form.Project

	project, err := bindProjectFormAndMapToProject(ctx, projectForm, c.mapper)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := c.service.SaveProject(project); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (c *project) DeleteProject(ctx *gin.Context) {
	id, err := getIdParamFromContext(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := c.service.DeleteProjectById(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (c *project) UpdateProject(ctx *gin.Context) {
	id, err := getIdParamFromContext(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	var projectForm *form.Project

	project, err := bindProjectFormAndMapToProject(ctx, projectForm, c.mapper)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := c.service.UpdateProjectById(id, project); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}
