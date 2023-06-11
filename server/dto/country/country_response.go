package countrydto

type CountryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"type: varchar(255)"`
}