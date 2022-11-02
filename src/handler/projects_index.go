package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getAllProjects(c *gin.Context) {
	projects, err := h.GetAllProjects()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"projects": projects.All(),
	})
}
