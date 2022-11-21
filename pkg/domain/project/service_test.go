package project

import (
	com "github.com/mihai-valentin/projects-monitor/pkg/com/project"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/entity"
	"github.com/mihai-valentin/projects-monitor/pkg/domain/project/repository"
	"testing"
)

func setUpTest() (expectedProjects entity.ProjectsList, projectService *Project) {
	expectedProjects = *entity.FakeProjectsList(3)
	r := repository.Mock()

	repositoryProjects := expectedProjects
	r.SetData(&repositoryProjects)

	m := NewMapper()

	projectService = NewService(m, r)

	return
}

func TestProject_GetAllProjects(t *testing.T) {
	expectedProjects, service := setUpTest()

	projects := service.GetAllProjects()

	t.Run("Get all projects and check projects count", func(t *testing.T) {
		if projects.Count() != expectedProjects.Count() {
			t.Fatal("Expected projects count doesn't match actual")
		}
	})

	t.Run("Get all projects and check each project in list", func(t *testing.T) {
		projectsSlice := projects.All()

		for i, project := range expectedProjects.All() {
			if projectsSlice[i].Id == project.Id {
				continue
			}

			t.Fatal("Unexpected project found in the list")
		}
	})
}

func TestProject_GetProjectById(t *testing.T) {
	expectedProjects, service := setUpTest()

	t.Run("Try to get a nonexistent project", func(t *testing.T) {
		project, err := service.GetProjectById(-1)

		if project != nil {
			t.Fatal("Unexpected project found")
		}

		if err == nil {
			t.Fatal("Expected error missing")
		}

		if err.Error() != "project not found" {
			t.Fatal("Wrong error missing")
		}
	})

	t.Run("Get an existing project", func(t *testing.T) {
		expectedProject := expectedProjects.All()[0]
		project, err := service.GetProjectById(expectedProject.Id)

		if err != nil {
			t.Fatal("Unexpected error")
		}

		if project == nil {
			t.Fatal("Nil project found")
		}

		if project.Id != expectedProject.Id {
			t.Fatal("Unexpected project found")
		}
	})
}

func TestProject_SaveProject(t *testing.T) {
	_, service := setUpTest()

	t.Run("Save project successfully", func(t *testing.T) {
		project := &com.Project{
			Name:        "test",
			Description: "test",
		}

		if err := service.SaveProject(project); err != nil {
			t.Fatal("Unexpected error")
		}

		projectsList := service.GetAllProjects()

		lastProject := projectsList.All()[projectsList.Count()-1]

		if lastProject.Name != project.Name || lastProject.Description != project.Description {
			t.Fatal("Project not found in storage")
		}
	})
}

func TestProject_UpdateProjectById(t *testing.T) {
	expectedProjects, service := setUpTest()

	t.Run("Try to update nonexistent project", func(t *testing.T) {
		err := service.UpdateProjectById(-1, &com.Project{})

		if err == nil {
			t.Fatal("Expected error missing")
		}

		if err.Error() != "project not found" {
			t.Fatal("Wrong error missing")
		}
	})

	t.Run("Update project successfully", func(t *testing.T) {
		project := expectedProjects.All()[0]
		updateData := &com.Project{
			Name:        "new name",
			Description: project.Description,
		}

		if err := service.UpdateProjectById(project.Id, updateData); err != nil {
			t.Fatal("Unexpected error")
		}

		updatedProject, _ := service.GetProjectById(project.Id)

		if updatedProject == nil || updatedProject.Name != updateData.Name {
			t.Fatal("Project was not updated in the storage")
		}
	})
}

func TestProject_DeleteProjectById(t *testing.T) {
	expectedProjects, service := setUpTest()

	t.Run("Try to delete nonexistent project", func(t *testing.T) {
		err := service.DeleteProjectById(-1)

		if err == nil {
			t.Fatal("Expected error missing")
		}

		if err.Error() != "project not found" {
			t.Fatal("Wrong error missing")
		}
	})

	t.Run("Delete project successfully", func(t *testing.T) {
		project := expectedProjects.All()[0]

		if err := service.DeleteProjectById(project.Id); err != nil {
			t.Fatal("Unexpected error")
		}

		deletedProject, _ := service.GetProjectById(project.Id)

		if deletedProject != nil {
			t.Fatal("Project was not removed from the storage")
		}
	})
}
