package main

import (
	projectDomain "github.com/mihai-valentin/projects-monitor/pkg/domain/project"
	transport "github.com/mihai-valentin/projects-monitor/pkg/transport/api"
	projectApi "github.com/mihai-valentin/projects-monitor/pkg/transport/api/project"
)

func main() {
	pd := projectDomain.New()
	pc := projectApi.New(pd)

	router := transport.NewRouter()
	router.RegisterController(pc)

	if err := router.Run(); err != nil {
		panic(err)
	}
}
