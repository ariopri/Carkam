package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type KampusSuccsesRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type KampusSucssesResponse struct {
	Message string `json:"message"`
}

type Review struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreatereviewRequest struct {
	Isian string `json:"isian"`
}

type CreatereviewResponse struct {
	Message string `json:"message"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterSuccessResponse struct {
	Message string `json:"message"`
}

type LoginSuccessResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type LogoutSuccessResponse struct {
	Message string `json:"message"`
}

type AuthErrorResponse struct {
	Error string `json:"error"`
}

var jwtKey = []byte("key")

type Claims struct {
	Username string
	jwt.StandardClaims
}

func (api *API) createreview(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	var req CreatereviewRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = api.reviewRepo.InsertcreatReview(req.Isian)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(CreatereviewResponse{Message: "Berhasil di tambahkan"})
}

func (api *API) review(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	var req Review
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = api.reviewRepo.InsertReview(req.Name, req.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(CreatereviewResponse{Message: "Berhasil di tambahkan"})
}

func (api *API) kampus(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	var req KampusSuccsesRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = api.kampusRepo.InsertKampus(req.Name, req.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(KampusSucssesResponse{Message: "Berhasil di tambahkan"})
}

func (api *API) register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = api.usersRepo.InsertUser(req.Username, req.Email, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(RegisterSuccessResponse{Message: "Berhasil di tambahkan"})
}

func (api *API) login(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := api.usersRepo.Login(user.Username, user.Password)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(AuthErrorResponse{Error: "Login Failed"})
		return
	}
	expirationTime := time.Now().Add(60 * time.Minute)

	claims := &Claims{
		Username: *res,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		Path:    "/",
	})
	json.NewEncoder(w).Encode(LoginSuccessResponse{Username: *res, Token: tokenString})
}

func (api *API) logout(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)

	token, err := req.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// return unauthorized ketika token kosong
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// return bad request ketika field token tidak ada
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if token.Value == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	c := http.Cookie{
		Name:   "token",
		MaxAge: -1,
	}
	http.SetCookie(w, &c)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUKSES LOGOUT"))
}
