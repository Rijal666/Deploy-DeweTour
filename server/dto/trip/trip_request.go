package tripdto

import "dewetour/models"

type CreateTripRequest struct {
	ID             int    `json:"id"`
	Title          string `json:"title" form:"title" binding:"required"`
	CountryId      int    `json:"country_id" binding:"required" form:"country_id"`	
	Accomodation   string `json:"accomodation" form:"accomodation" binding:"required"`
	Transportation string `json:"transportation" form:"transportation" binding:"required"`
	Eat            string `json:"eat" form:"eat" binding:"required"`
	Day            int    `json:"day" form:"day" binding:"required"`
	Night          int    `json:"night" form:"night" binding:"required"`
	DateTrip       string `json:"date_trip" form:"date_trip" binding:"required"`
	Price          int    `json:"price" form:"price" binding:"required"`
	Quota          int    `json:"quota" form:"quota" binding:"required"`
	Description    string `json:"description" form:"description" binding:"required"`
	Image          string `json:"image" form:"image" binding:"required"`
}

type Update1TripRequest struct {
	ID        int    `json:"id"`
	Title     string `json:"title" form:"title"`
	CountryId int    `json:"country_id" form:"country_id"`
	// Country        models.CountryResponse `json:"country"`
	Accomodation   string `json:"accomodation" form:"accomodation"`
	Transportation string `json:"transportation" form:"transportation"`
	Eat            string `json:"eat" form:"eat"`
	Day            int    `json:"day" form:"day"`
	Night          int `json:"night" form:"night"`
	DateTrip       string `json:"date_trip" form:"date_trip"`
	Price          int    `json:"price" form:"price"`
	Quota          int    `json:"quota" form:"quota"`
	Description    string `json:"description" form:"description"`
	Image          string `json:"image" form:"image"`
}

type TripResponse struct {
	ID        int `json:"id"`
	Title string `json:"title"`
	CountryID int `json:"country_id"`
	Country   models.CountryResponse `json:"country"`
	Accomodation   string `json:"accomodation" `
	Transportation string `json:"transportation" `
	Eat            string `json:"eat" `
	Day            int `json:"dat" `
	Night          int `json:"night" `
	DateTrip       string `json:"date_trip" `
	Price          int `json:"price" `
	Quota          int `json:"quota" `
	Description    string `json:"description" `
	Image          string `json:"image"`
}