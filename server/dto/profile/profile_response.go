package profiledto

import "dewetour/models"

type ProfileResponse struct {
	ID     int `json:"id" gorm:"primary_key:auto_increment"`
	UserID int `json:"user_id"`
	User   models.UsersProfileResponse`json:"user"`
	Image string `json:"image" gorm:"type: varchar(255)"`
}