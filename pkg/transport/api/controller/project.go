package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mihai-valentin/projects-monitor/pkg/com"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"
	"github.com/mihai-valentin/projects-monitor/pkg/transport/api/form"
	"net/http"
	"strconv"
)

type Project struct {
	service com.Project
}

func NewProjectController(s com.Project) *Project {
	return &Project{s}
}

func (c *Project) RegisterRoutes(gr *gin.Engine) {
	projects := gr.Group("/projects")
	{
		projects.GET("/:id", c.getProjectById)
		projects.PUT("/:id", c.updateProject)
		projects.DELETE("/:id", c.deleteProject)
	}
}

func (c *Project) getProjectById(ctx *gin.Context) {
	id, err := c.getIdParamFromContext(ctx)

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

	ctx.JSON(http.StatusOK, p)
}

func (c *Project) deleteProject(ctx *gin.Context) {
	id, err := c.getIdParamFromContext(ctx)

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

func (c *Project) updateProject(ctx *gin.Context) {
	id, err := c.getIdParamFromContext(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

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

	if err := c.service.UpdateProjectById(id, p); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}

func (c *Project) getIdParamFromContext(ctx *gin.Context) (int, error) {
	idRaw := ctx.Param("id")
	id, err := strconv.Atoi(idRaw)

	if err != nil {
		return 0, errors.New("path parameter id must be an int")
	}

	return id, nil
}
