package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mihai-valentin/projects-monitor/pkg/com"
	"net/http"
	"strconv"
)

type ProjectsPaginatedList struct {
	service com.Project
}

func NewProjectsPaginatedList(s com.Project) *ProjectsPaginatedList {
	return &ProjectsPaginatedList{s}
}

func (c *ProjectsPaginatedList) RegisterRoutes(gr *gin.Engine) {
	gr.GET("/catalog/projects", c.getAllProjectsPaginated)
}

func (c *ProjectsPaginatedList) getAllProjectsPaginated(ctx *gin.Context) {
	pageRaw := ctx.Query("page")
	page, _ := strconv.Atoi(pageRaw)

	if page == 0 {
		page = 1
	}

	perPageRaw := ctx.Query("per_page")
	perPage, _ := strconv.Atoi(perPageRaw)

	if perPage == 0 {
		perPage = 3
	}

	ppl := c.service.GetAllProjectsPaginated(page, perPage)

	ctx.JSON(http.StatusOK, ppl)
}
