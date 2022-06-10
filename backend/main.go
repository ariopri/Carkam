package main

import (
	"github.com/rg-km/final-project-engineering-66/api"

	"github.com/ariopri/rg-km/final-project-engineering-66//db"
	"github.com/ariopri/rg-km/final-project-engineering-66//repository"
)

func main() {
	db := &db.CsvDB{}
	usersRepo := repository.NewUserRepository(db)
	productsRepo := repository.NewProductRepository(db)

	mainAPI := api.NewAPI(usersRepo, productsRepo)
	mainAPI.Start()

}
