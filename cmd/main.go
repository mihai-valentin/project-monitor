package main

import (
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project"
	transport "github.com/mihai-valentin/projects-monitor/pkg/transport/api"
	"github.com/mihai-valentin/projects-monitor/pkg/transport/api/controller"
)

func main() {
	r := transport.NewRouter()

	pd := project.New()

	pc := controller.NewProjectController(pd)
	r.RegisterController(pc)

	plc := controller.NewProjectsListController(pd)
	r.RegisterController(plc)

	pplc := controller.NewProjectsPaginatedList(pd)
	r.RegisterController(pplc)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
