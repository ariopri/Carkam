package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	gubrak "github.com/novalagung/gubrak/v2"
)

type CustomMux struct {
	http.ServeMux
	middlewares []func(next http.Handler) http.Handler
}

func (c *CustomMux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux

	for _, next := range c.middlewares {
		current = next(current)
	}

	current.ServeHTTP(w, r)
}

type M map[string]interface{}

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"Username"`
	Email    string `json:"Email"`
}

var APPLICATION_NAME = "Simple JWT App"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of camp2022")

func main() {
	mux := new(CustomMux)
	mux.RegisterMiddleware(MiddlewareJWTAuthorization) // /dashboard -> MiddlewareJWTAuthorization -> HandlerDashboard

	mux.HandleFunc("/register", HandlerRegister)
	mux.HandleFunc("/login", HandlerLogin)
	mux.HandleFunc("/logout", HandlerLogout)

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":9090"

	fmt.Println("Starting server at", server.Addr)
	server.ListenAndServe()
}

func HandlerRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Unsupported http method", http.StatusBadRequest)
		return
	}
}

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Unsupported http method", http.StatusBadRequest)
		return
	}

	// Basic validation
	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "Invalid username or password basic", http.StatusBadRequest)
		return
	}

	// Authenticate user
	ok, userInfo := authenticateUser(username, password)
	if !ok {
		http.Error(w, "Invalid username or password data users.json", http.StatusBadRequest)
		return
	}

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Username: userInfo["username"].(string),
		Email:    userInfo["email"].(string),
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenString, _ := json.Marshal(M{"token": signedToken})
	w.Write([]byte(tokenString))
}

func HandlerLogout(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Unsupported http method", http.StatusBadRequest)
		return
	}
	w.Write([]byte("Logout success"))
}

func authenticateUser(username, password string) (bool, M) {
	dbPath, err := filepath.Abs("./data/users.json")
	if err != nil {
		return false, nil
	}

	buf, _ := ioutil.ReadFile(dbPath)

	data := make([]M, 0)
	errUnmarshal := json.Unmarshal(buf, &data)
	if errUnmarshal != nil {
		return false, nil
	}

	res := gubrak.From(data).Find(func(each M) bool {
		return each["username"] == username && each["password"] == password
	}).Result()

	if res != nil {
		resM := res.(M)
		delete(resM, "password")
		return true, resM
	}

	return false, nil
}

func MiddlewareJWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		authorizationHeader := r.Header.Get("Authorization") // -H "Authorization: Bearer ...."
		if !strings.Contains(authorizationHeader, "Bearer") {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1) // ....

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("signing method invalid")
			}

			return JWT_SIGNATURE_KEY, nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(context.Background(), "userInfo", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// curl -X POST --user aditira:aditira123 http://localhost:8080/login
// {"token":"xxxxxx.yyyyyy.zzzzzzz"}

// curl -X GET --header "Authorization: Bearer xxxxxx.yyyyyy.zzzzzzz" http://localhost:8080/dashboard
