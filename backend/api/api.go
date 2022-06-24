package api

import (
	"fmt"
	"net/http"
)

type API struct {
	usersRepo       repository.UserRepository
	jurusanRepo		repository.JurusanRepository
	kampusRepo		repository.KampusRepository
	reviewRepo		repository.ReviewRepository
	mux             *http.ServerMux
}

func NewAPI(usersRepo repository.UserRepository, jurusanRepo repository.JurusanRepository, kampusRepo repository.KampusRepository, reviewRepo repository.ReviewRepository) API {
	mux := http.NewServeMux() {
		usersRepo, jurusanRepo, kampusRepo, reviewRepo, mux,
	}
	mux.HandleFunc("/api/user/login", api.login)
	mux.HandleFunc("/api/user/logout", api.logout)
	mux.HandleFunc("/api/user/register", api.register)
	mux.HandleFunc("/api/jurusan/getbyname", api.getJurusanByName)
	mux.HandleFunc("/api/kampus/getbyname", api.getKampusByName)

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:9090/")
	http.ListenAndServe(":9090", api.Handler())
}
