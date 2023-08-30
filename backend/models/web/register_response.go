package web

import "time"

type RegisterResponse struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}
