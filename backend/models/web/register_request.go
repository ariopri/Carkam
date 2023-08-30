package web

import "time"

type RegisterRequest struct {
	FirstName string    `json:"firstname" validate:"required,min=1,max=50"`
	LastName  string    `json:"lastname" validate:"required,min=1,max=50"`
	UserName  string    `json:"username" validate:"required,min=1,max=50,lowercase"`
	Email     string    `json:"email" validate:"required,min=1,max=50,email"`
	Phone     string    `json:"phone" validate:"required,min=1,max=50"`
	Password  string    `json:"password" validate:"required,min=1,max=50"`
	CreatedAt time.Time `json:"created_at"`
}
