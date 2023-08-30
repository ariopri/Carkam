package web

type LoginResponse struct {
	ID        string `json:"id"`
	UserName  string `json:"username"`
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
	Role      string `json:"role"`
}
