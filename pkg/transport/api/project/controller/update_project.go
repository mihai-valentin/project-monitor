package project

import (
	"github.com/gin-gonic/gin"
	"github.com/mihai-valentin/projects-monitor/pkg/transport/api/project/form"
	"net/http"
)

func (c *Controller) UpdateProject(ctx *gin.Context) {
	id, err := c.getIdParamFromContext(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	var projectForm *form.Project

	project, err := c.bindProjectFormAndMapToProject(projectForm, ctx)

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
