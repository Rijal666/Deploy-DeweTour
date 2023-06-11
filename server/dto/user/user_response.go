package userdto

type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Image    string `json:"image" form:"image"`
}