package web

import "time"

type UpdateUserRequest struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name" validate:"required,min=1,max=50"`
	LastName  string    `json:"last_name" validate:"required,min=1,max=50"`
	UserName  string    `json:"username" validate:"required,min=1,max=50,lowercase"`
	Email     string    `json:"email" validate:"required,min=1,max=50,email"`
	Phone     string    `json:"phone" validate:"required,min=1,max=50"`
	Kota      string    `json:"kota" validate:"required,min=1,max=50"`
	Avatar    string    `json:"avatar" validate:"required,min=1,max=50"`
	UpdateAt  time.Time `json:"update_at"`
}
