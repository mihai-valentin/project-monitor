package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/action"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/mapper"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/repository"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/service"
)

func RegisterDomain(gr *gin.Engine) {
	r := repository.New()
	s := service.New(r)
	fm := mapper.NewProjectFormMapper()
	pm := mapper.NewProjectJsonMapper()
	plm := mapper.NewProjectListJsonMapper()

	action.NewGetAllProjectsAction(plm, s).LoadRoutes(gr)
	action.NewGetProjectAction(pm, s).LoadRoutes(gr)
	action.NewCreateProjectAction(fm, s).LoadRoutes(gr)
	action.NewUpdateProjectAction(fm, s).LoadRoutes(gr)
	action.NewDeleteProjectAction(s).LoadRoutes(gr)
}
