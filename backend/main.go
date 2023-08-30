package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"github.com/rg-km/final-project-engineering-66/config"
	"github.com/rg-km/final-project-engineering-66/controllers"
	"github.com/rg-km/final-project-engineering-66/repository"
	"github.com/rg-km/final-project-engineering-66/route"
	"github.com/rg-km/final-project-engineering-66/service"
)

func main() {
	loadConfig, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	db, err := config.ConnectionDB(&loadConfig)
	if err != nil {
		panic(err)
	}
	validate := validator.New()

	userRepository := repository.NewUsersRepository(db)
	userService := service.NewUserServiceImpl(userRepository, db, validate)
	userController := controllers.NewUserController(userService)

	router := route.NewRouter(userRepository, userController)

	err = router.Run(":8080")
	if err != nil {
		return
	}
}
