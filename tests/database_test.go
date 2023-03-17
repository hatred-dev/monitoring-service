package tests

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"monitoring-service/database"
	"testing"
)

func TestCreateProject(t *testing.T) {
	createDB()
	defer dropDB()
	ctx := context.Background()
	queries := database.New(dbConnection)
	project, err := queries.CreateProject(ctx, database.CreateProjectParams{
		ProjectName: "test_project",
		Active: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(project)
}

func TestCreateMultipleProjects(t *testing.T) {
	createDB()
	ctx := context.Background()
	queries := database.New(dbConnection)
	for i := 1; i <= 100; i++ {
		_, err := queries.CreateProject(ctx, database.CreateProjectParams{
			ProjectName: fmt.Sprintf("%dproject", i),
			Active: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
		})
		if err != nil {
			t.Error(err)
		}
	}
	projects, err := queries.GetProjects(ctx)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(projects[0:4])
	assert.Equal(t, len(projects), 100)
}

func TestProjectsDelete(t *testing.T) {
	defer dropDB()
	ctx := context.Background()
	queries := database.New(dbConnection)
	for i := 1; i <= 100; i++ {
		err := queries.DeleteProject(ctx, fmt.Sprintf("%dproject", i))
		if err != nil {
			t.Error(err)
		}
	}
	projects, err := queries.GetProjects(ctx)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, len(projects), 0)
}

func TestInsertIP(t *testing.T) {
	createDB()
	defer dropDB()

	ctx := context.Background()
	queries := database.New(dbConnection)
	project, err := queries.CreateProject(ctx, database.CreateProjectParams{
		ProjectName: "test_project",
		Active: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
	})
	if err != nil {
		t.Error(err)
	}
	ip, err := queries.CreateIP(ctx, database.CreateIPParams{
		ProjectName: project.ProjectName,
		Ip:          "1.1.1.1",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ip)
}

func TestCascadeDelete(t *testing.T) {
	createDB()
	defer dropDB()

	ctx := context.Background()
	queries := database.New(dbConnection)
	project, err := queries.CreateProject(ctx, database.CreateProjectParams{
		ProjectName: "test_project",
		Active: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
	})
	if err != nil {
		t.Error(err)
	}
	ip, err := queries.CreateIP(ctx, database.CreateIPParams{
		ProjectName: project.ProjectName,
		Ip:          "1.1.1.1",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ip)
	err = queries.DeleteProject(ctx, project.ProjectName)
	ips, err := queries.GetAllIPs(ctx)
	assert.Equal(t, len(ips), 0)
}
