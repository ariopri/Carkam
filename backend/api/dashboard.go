package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Kampus struct {
	id         int    `json:"id"`
	name       string `json:"name"`
	address    string `json:"address"`
	phone      int    `json:"phone"`
	email      string `json:"email"`
	website    string `json:"website"`
	akreditasi string `json:"akreditasi"`
}

type DashboardSuccessResponse struct {
	kampus []Kampus `json:"kampus"`
}

func (api *API) getDashboard(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)
	username := req.URL.Query().Get("username")

	Kampus, err := api.repository.GetKampus(username)
	defer func() {
		if err != nil {
			encoder.Encode(AdminErrorResponse{Error: err.Error()})
			return
		}
	}()
	if err != nil {
		encoder.Encode(AdminErrorResponse{Error: err.Error()})
		return
	}
	encoder.Encode(DashboardSuccessResponse{kampus: Kampus})
	w.WriteHeader(http.StatusOK)
	respons := DashboardSuccessResponse{kampus: Kampus}
	for _, kampus := range respons.kampus {
		fmt.Println(kampus.name)
	}
	encoder.Encode(respons)
}
