package handler

import (
	"github.com/gin-gonic/gin"
	"projects-monitor/src/entity"
)

type Service interface {
	GetAllProjects() (*entity.ProjectsList, error)
}

type Handler struct {
	Service
}

func New(s Service) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) InitGinRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/projects", h.getAllProjects)

	return r
}
