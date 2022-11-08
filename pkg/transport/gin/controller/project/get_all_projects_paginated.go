package project

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const defaultPerPage = 10
const defaultPage = 1

func (c *Project) GetAllProjectsPaginated(ctx *gin.Context) {
	perPageRaw := ctx.Query("per_page")
	perPage, err := strconv.Atoi(perPageRaw)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": "query parameter per_page must be an int",
		})
	}

	if perPage <= 0 {
		perPage = defaultPerPage
	}

	pageRaw := ctx.Query("page")
	page, err := strconv.Atoi(pageRaw)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": "query parameter page must be an int",
		})
	}

	if page <= 0 {
		page = defaultPage
	}

	paginatedProjects := c.service.GetAllProjectsPaginated(page, perPage)

	ctx.JSON(http.StatusOK, gin.H{
		"items":        paginatedProjects.Items,
		"total_pages":  paginatedProjects.Pagination.TotalPages,
		"current_page": paginatedProjects.Pagination.CurrentPage,
		"per_page":     paginatedProjects.Pagination.PerPage,
	})
}
