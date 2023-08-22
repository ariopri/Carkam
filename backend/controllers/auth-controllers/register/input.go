package registerAuth

type InputRegister struct {
	FirstName string `json:"firstname" validate:"required,min=1,max=50"`
	LastName  string `json:"lastname" validate:"required,min=1,max=50"`
	UserName  string `json:"username" validate:"required,min=1,max=50,lowercase"`
	Email     string `json:"email" validate:"required,min=1,max=50,email"`
	Password  string `json:"password" validate:"required,min=1,max=50"`
}
