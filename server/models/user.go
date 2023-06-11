package models

type User struct {
	ID          int                   `json:"id" gorm:"primary_key:auto_increment"`
	IsAdmin     bool                  `json:"is_admin" gorm:"type: bool"`
	Fullname    string                `json:"fullname" gorm:"type: varchar(255)"`
	Email       string                `json:"email" gorm:"type: varchar(255)"`
	Password    string                `json:"password" gorm:"type: varchar(255)"`
	Phone       string                `json:"phone" gorm:"type: varchar(255)"`
	Address     string                `json:"address" gorm:"type: varchar(255)"`
	Image       string                `json:"image" gorm:"type: varchar(255)"`
	Transaction []TransactionResponse `json:"transaction"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
