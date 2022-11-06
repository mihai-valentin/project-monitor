package action

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteService interface {
	DeleteProjectById(id int) error
}

type DeleteProject struct {
	service DeleteService
}

func NewDeleteProjectAction(s DeleteService) *DeleteProject {
	return &DeleteProject{s}
}

func (a *DeleteProject) LoadRoutes(gr *gin.Engine) {
	gr.DELETE("/projects/:id")
}

func (a *DeleteProject) deleteProject(ctx *gin.Context) {
	idRaw := ctx.Param("id")
	id, err := strconv.Atoi(idRaw)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": "path parameter id must be an int",
		})

		return
	}

	if err := a.service.DeleteProjectById(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
