package repository

import (
	"database/sql"
	"projects-monitor/src/entity"
)

type Project struct {
	dbCon *sql.DB
}

func newProjectsRepository(dbCon *sql.DB) *Project {
	return &Project{dbCon: dbCon}
}

func (p *Project) GetAll() (*entity.ProjectsList, error) {
	rows, _ := p.dbCon.Query("SELECT * FROM projects")
	defer rows.Close()

	projects := entity.NewProjectsList()

	for rows.Next() {
		project, err := p.mapProject(rows)
		if err != nil {
			return nil, err
		}

		projects.Add(project)
	}

	return projects, nil
}

func (p *Project) mapProject(rows *sql.Rows) (*entity.Project, error) {
	var id int
	var name string

	if err := rows.Scan(&id, &name); err != nil {
		return nil, err
	}

	return entity.NewProject(id, name), nil
}
