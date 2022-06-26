package main

import (
	"database/sql"

	"github.com/rg-km/final-project-engineering-66/api"
	"github.com/rg-km/final-project-engineering-66/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./CARKAMD.db")
	if err != nil {
		panic(err)
	}
	usersRepo := repository.NewUserRepository(db)
	jurusanRepo := repository.NewJurusanRepository(db)
	kampusRepo := repository.NewKampusRepository(db)
	reviewRepo := repository.NewReviewRepository(db)

	api := api.NewAPI(*usersRepo, *kampusRepo, *jurusanRepo, *reviewRepo)
	api.Start()
}
