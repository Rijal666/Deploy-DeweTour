package authdto

type AuthRequset struct {
	IsAdmin  bool   `json:"is_admin"`
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Image    string `json:"image"`
}

type LoginRequest struct {
	IsAdmin  bool   `json:"is_admin"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}