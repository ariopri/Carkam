package web

type LoginResponse struct {
	ID        string `json:"id"`
	Role      string `json:"role"`
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}
