package main

import (
	project "github.com/mihai-valentin/projects-monitor/pkg/domain/project/provider"
	ginTransport "github.com/mihai-valentin/projects-monitor/pkg/transport/gin/provider"
)

func main() {
	projectDomain := project.RegisterDomain()

	transport := ginTransport.New()
	transport.LoadProjectDomain(projectDomain)

	if err := transport.Run(); err != nil {
		panic(err)
	}
}
