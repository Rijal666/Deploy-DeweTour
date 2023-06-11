package authdto

type LoginResponse struct {
	ID      int    `json:"id"`
	IsAdmin bool   `json:"is_admin"`
	Email   string `json:"email"`
	Token   string `json:"token"`
}