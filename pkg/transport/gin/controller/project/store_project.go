package project

import (
	"github.com/gin-gonic/gin"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"
	"net/http"
)

func (c *Project) StoreProject(ctx *gin.Context) {
	var projectForm *entity.ProjectForm

	project, err := c.bindProjectFormAndMapToProject(projectForm, ctx)

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