package main

import (
	"database/sql"
	"fmt"

	"github.com/rg-km/final-project-engineering-66/api"
	"github.com/rg-km/final-project-engineering-66/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./carkam.db")
	if err != nil {
		panic(err)
	}
	usersRepo := repository.NewUserRepository(db)
	jurusanRepo := repository.NewJurusanRepository(db)
	kampusRepo := repository.NewKampusRepository(db)
	reviewRepo := repository.NewReviewRepository(db)
	fmt.Println("starting web server at http://localhost:8080")
	api := api.NewAPI(*usersRepo, *kampusRepo, *jurusanRepo, *reviewRepo)
	api.Start()
}
