package project

import "github.com/mihai-valentin/projects-monitor/pkg/domain/project/repository"

func New() *Project {
	r := repository.New()
	m := NewMapper()

	return NewService(m, r)
}
