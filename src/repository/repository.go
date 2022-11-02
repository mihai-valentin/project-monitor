package repository

import "database/sql"

type Repository struct {
	*Project
}

func New(dbCon *sql.DB) *Repository {
	return &Repository{
		Project: newProjectsRepository(dbCon),
	}
}
