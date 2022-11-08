package project

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Project) GetProjectById(ctx *gin.Context) {
	id, err := c.getIdParamFromContext(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	project, err := c.service.GetProjectById(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, c.mapper.MapProjectToJson(project))
}
