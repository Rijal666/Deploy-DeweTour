package profiledto

type ProfileRequest struct {
	ID     int    `json:"id" gorm:"primary_key:auto_increment" validate:"required"`
	UserID int    `json:"user_id"`
	Image  string `json:"image" gorm:"type: varchar(255)" validate:"required"`
}