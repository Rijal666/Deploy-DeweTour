package countrydto

type CreateCountryRequest struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name" gorm:"type: varchar(255)" validate:"required"`
}

type UpdateCountryRequset struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}