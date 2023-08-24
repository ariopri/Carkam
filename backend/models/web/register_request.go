package web

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required,min=1,max=50"`
	LastName  string `json:"last_name" validate:"required,min=1,max=50"`
	UserName  string `json:"username" validate:"required,min=1,max=50,lowercase"`
	Email     string `json:"email" validate:"required,min=1,max=50,email"`
	Phone     string `json:"phone" validate:"required,min=1,max=50"`
	Password  string `json:"password" validate:"required,min=1,max=50"`
}
