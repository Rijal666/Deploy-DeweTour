package userdto

type CreateUserRequest struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Image    string `json:"image" form:"image"`
}

type UpdateUserRequest struct {
	IsAdmin  bool   `json:"is_admin"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Image    string `json:"image" `
}