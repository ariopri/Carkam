package api

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/go-delve/delve/service/api"
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
		mux.Hadnle("/api/users/login", api.GET(api.AuthMiddleware(http.HandlerFunc(api.HandlerLogin))))
		mux.Hadnle("/api/users/register", api.POST(api.AuthMiddleware(http.HandlerFunc(api.HandlerRegister))))
		mux.Hadnle("/api/users/logout", api.GET(api.AuthMiddleware(http.HandlerFunc(api.HandlerLogout))))
		mux.Handle("api/jurusan/name", api.GET(api.AuthMiddleware(http.HandlerFunc(api.HandlerJurusanName))))
		mux.Handle("api/kampus/name", api.GET(api.AuthMiddleware(http.HandlerFunc(api.HandlerKampusName))))	
		return api
	}
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:9090/")
	http.ListenAndServe(":9090", api.Handler())
}
