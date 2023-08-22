package create

import "time"

type InputCreateStudent struct {
	FirstName string `json:"first_name" validate:"required,min=1,max=50"`
	LastName  string `json:"last_name" validate:"required,min=1,max=50"`
	Email     string `json:"email" validate:"required,min=1,max=50,email"`
	Phone     string `json:"phone" validate:"required,min=1,max=50"`
	Kota      string `json:"kota" validate:"required,min=1,max=50"`
	Avatar    string `json:"avatar" validate:"required,min=1,max=50"`
	CreatedAt time.Time
}
