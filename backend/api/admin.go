package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type AdminErrorResponse struct {
	Error string `json:"error"`
}

type ProductKampus struct {
	id         int    `json:"id"`
	name       string `json:"name"`
	address    string `json:"address"`
	phone      int    `json:"phone"`
	email      string `json:"email"`
	website    string `json:"website"`
	akreditasi string `json:"akreditasi"`
}

type AdminResponse struct {
	Kampus []repository.kampus `json:"kampus"`
}

func (api *API) getDashboard(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	kampusName := req.URL.Query().Get("kampus_name")
	var startPeriod time.Time
	var endPeriod time.Time

	startPeriod, err := time.Parse("2022-06-27", req.URL.Query().Get("start_period"))
	endPeriod, _ = time.Parse("2022-06-28", req.URL.Query().Get("end_period"))

	getKampusRequest := repository.GetKampusRequest{
		KampusName:  kampusName,
		StartPeriod: startPeriod,
		EndPeriod:   endPeriod,
	}
	if req.URL.Query().Get("start_period") == "" {
		getKampusRequest.StartPeriod = time.Now()
	}
	if req.URL.Query().Get("end_period") == "" {
		getKampusRequest.EndPeriod = time.Now()
	}

	encoder := json.NewEncoder(w)

	kampus, err := api.repository.GetKampus(getKampusRequest)
	if err != nil {
		encoder.Encode(AdminErrorResponse{Error: err.Error()})
		return
	}

	encoder.Encode(AdminResponse{Kampus: kampus})
	w.WriteHeader(http.StatusOK)
}
