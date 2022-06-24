package main

import (
	"github.com/rg-km/final-project-engineering-66/api"

	"github.com/ariopri/rg-km/final-project-engineering-66/db"
	"github.com/ariopri/rg-km/final-project-engineering-66/repository"
)

func main() {
	db := &db.CsvDB{}
	usersRepo := repository.NewUserRepository(db)
	jurusanRepo := repository.NewJurusanRepository(db)
	kampusRepo := repository.NewKampusRepository(db)
	reviewRepo := repository.NewReviewRepository(db)
	api := api.NewAPI(usersRepo, jurusanRepo, kampusRepo, reviewRepo)
	api.Start()
}
