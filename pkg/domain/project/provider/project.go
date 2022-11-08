package provider

import (
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/repository"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/service"
)

func RegisterDomain() *service.Project {
	r := repository.New()

	return service.New(r)
}
