package main

import (
	projectDomain "github.com/mihai-valentin/projects-monitor/pkg/domain/project"
	"github.com/mihai-valentin/projects-monitor/pkg/transport/api"
	projectController "github.com/mihai-valentin/projects-monitor/pkg/transport/api/project"
)

func main() {
	pd := projectDomain.New()
	pc := projectController.New(pd)

	transport := api.New()
	transport.RegisterController(pc)

	if err := transport.Run(); err != nil {
		panic(err)
	}
}
