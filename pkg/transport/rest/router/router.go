package router

import "github.com/gin-gonic/gin"

type Middleware interface {
	CORS(c *gin.Context)
}

type Router struct {
	middleware Middleware
}

func New(m Middleware) *Router {
	return &Router{
		middleware: m,
	}
}

func (r *Router) InitGinRouter() *gin.Engine {
	gr := gin.Default()

	gr.Use(r.middleware.CORS)

	return gr
}
