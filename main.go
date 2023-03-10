package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"monitoring-service/database"
)

func run() error {
	ctx := context.Background()
	connStr := "postgresql://postgres:postgres@0.0.0.0:5432/projects?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	queries := database.New(db)
	newProject, err := queries.CreateProject(ctx, database.CreateProjectParams{
		ProjectName: "test_project",
		Active: sql.NullBool{
			Bool:  false,
			Valid: true,
		},
	})
	if err != nil {
		return err
	}
	fmt.Println(newProject)
	return nil
}
func main() {
	//pingCmd := exec.Command("ping", "-c1", "-W", "15", "1.1.1.1")
	//pingRes, err := pingCmd.Output()
	//if err != nil {
	//	panic(err)
	//}
	//if strings.Contains(string(pingRes), "0 received") {
	//	fmt.Println()
	//}
	//fmt.Println(string(pingRes))
	run()
}
