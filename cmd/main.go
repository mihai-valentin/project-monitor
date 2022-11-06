package main

import (
	project "github.com/mihai-valentin/projects-monitor/pkg/domain/project/provider"
	transport "github.com/mihai-valentin/projects-monitor/pkg/transport/rest/provider"
)

func main() {
	gr := transport.RegisterTransport()
	project.RegisterDomain(gr)

	if err := gr.Run(); err != nil {
		panic(err)
	}
}
