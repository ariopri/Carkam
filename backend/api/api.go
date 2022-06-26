package api

import (
	"fmt"
	"net/http"

	"github.com/rg-km/final-project-engineering-66/repository"
)

type API struct {
	usersRepo   repository.UserRepository
	kampusRepo  repository.KampusRepository
	jurusanRepo repository.JurusanRepository
	reviewRepo  repository.ReviewRepository
	mux         *http.ServeMux
}

func NewAPI(usersRepo repository.UserRepository, kampusRepo repository.KampusRepository, jurusanRepo repository.JurusanRepository, reviewRepo repository.ReviewRepository) *API {
	mux := http.NewServeMux()
	api := API{
		usersRepo, kampusRepo, jurusanRepo, reviewRepo, mux,
	}

	mux.HandleFunc("/api/review", api.review)
	mux.HandleFunc("api/review", api.createreview)
	mux.HandleFunc("api/kampus", api.kampus)
	mux.HandleFunc("/api/register", api.register)
	mux.HandleFunc("/api/login", api.login)
	mux.HandleFunc("/api/logout", api.logout)

	return &api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8081", api.Handler())
}
