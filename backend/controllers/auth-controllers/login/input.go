package loginAuth

type InputLogin struct {
	UserName string `json:"username" validate:"required,min=1,max=50,lowercase"`
	Password string `json:"password" validate:"required,min=1,max=50"`
}
