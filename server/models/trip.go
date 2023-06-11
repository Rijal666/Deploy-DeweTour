package models

type Trip struct {
	ID             int             `json:"id"`
	Title          string          `json:"title" form:"title" validate:"required" gorm:"varchar(255)"`
	CountryId      int             `json:"country_id" validate:"required" form:"country_id"`
	Country        CountryResponse `json:"country" form:"country" gorm:"foreignKey: CountryId"`
	Accomodation   string          `json:"accomodation" form:"accomodation" validate:"required" gorm:"varchar(255)"`
	Transportation string          `json:"transportation" form:"transportation" validate:"required" gorm:"varchar(255)"`
	Eat            string          `json:"eat" form:"eat" gorm:"varchar(255)"`
	Day            int             `json:"day" form:"day" validate:"required" gorm:"varchar(255)"`
	Night          int             `json:"night" form:"night" validate:"required"`
	DateTrip       string          `json:"date_trip" form:"date_trip" validate:"required"`
	Price          int             `json:"price" form:"price" validate:"required"`
	Quota          int             `json:"quota" form:"quota" validate:"required"`
	Description    string          `json:"description" form:"description" validate:"required"`
	Image          string          `json:"image" form:"image" validate:"required" gorm:"varchar(255)"`
}

type TripResponse struct {
	ID             int             `json:"id"`
	Title          string          `json:"title"`
	CountryID      int             `json:"-"`
	Country        CountryResponse `json:"country"`
	Accomodation   string          `json:"accomodation"`
	Transportation string          `json:"transportation"`
	Eat            string          `json:"eat"`
	Day            int             `json:"day"`
	Night          int             `json:"night"`
	DateTrip       string          `json:"datetrip"`
	Price          int             `json:"price"`
	Quota          int             `json:"quota"`
	Description    string          `json:"description"`
	Image          string          `json:"image"`
}

func (TripResponse) TableName() string {
	return "trips"
}