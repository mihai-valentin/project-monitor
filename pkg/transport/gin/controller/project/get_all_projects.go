package project

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Project) GetAllProjects(ctx *gin.Context) {
	allProjects := c.service.GetAllProjects()
	projects := c.mapper.MapProjectsListToJson(allProjects)

	ctx.JSON(http.StatusOK, projects)
}
