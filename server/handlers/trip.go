package handlers

import (
	"context"
	resultdto "dewetour/dto/result"
	tripdto "dewetour/dto/trip"
	"dewetour/models"
	"dewetour/repositories"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

var path_file = "http://localhost:5000/uploads/"

type handlerTrip struct {
	TripRepository repositories.TripRepository
	CountryRepository repositories.CountryRepository
}

func HandlerTrip(TripRepository repositories.TripRepository, CountryRepository repositories.CountryRepository) *handlerTrip {
	return &handlerTrip{TripRepository, CountryRepository}
}

func(h *handlerTrip) FindTrips(c *gin.Context) {

		trips, err := h.TripRepository.FindTrips()
		
		if err != nil {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		// for i, p := range trips {
		// 	trips[i].Image = path_file + p.Image
		// } 

		if len(trips) > 0 {
			c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Data for all users was successfully obtained", Data: convertResponseTrips(trips)})
			return
		} else {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: "No record found"})
			return
		}
}

func (h *handlerTrip) GetTrip(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

		trip, err := h.TripRepository.GetTrip(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
				Status: http.StatusBadRequest, Message: err.Error()})
				return
		}

		// trip.Image = path_file + trip.Image

		c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Trip data successfully obtained", Data: trip})
		
}
func (h *handlerTrip) CreateTrip(c *gin.Context) {
	c.Header("Content-Type", "multipart/form-data")

	
	UserLogin := c.MustGet("userLogin")
	userAdmin := UserLogin.(jwt.MapClaims)["is_admin"].(bool)

	if userAdmin {
		dataFile := c.MustGet("dataFile").(string)
	fmt.Println("this is data file", dataFile)
		CountryId, _ := strconv.Atoi(c.Request.FormValue("country_id"))
		Day, _ := strconv.Atoi(c.Request.FormValue("day"))
		Night, _ := strconv.Atoi(c.Request.FormValue("night"))
		Price, _ := strconv.Atoi(c.Request.FormValue("price"))
		Quota, _ := strconv.Atoi(c.Request.FormValue("quota"))

		request := tripdto.CreateTripRequest{
			Title: c.Request.FormValue("title"),
			CountryId: CountryId,
			Accomodation: c.Request.FormValue("accomodation"),
			Transportation: c.Request.FormValue("transportation"),
			Eat: c.Request.FormValue("eat"),
			Day: Day,
			Night: Night,
			DateTrip: c.Request.FormValue("date_trip"),
			Price: Price,
			Quota: Quota,
			Description: c.Request.FormValue("description"),
			Image: dataFile,
		}
		
		
		// if err := c.Bind(request); err != nil{
		// 	c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		// 	return
		// }
		
		validation := validator.New()
		err := validation.Struct(request)
		
		if err != nil {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}
		
		var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")
		
	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "uploads"})

	if err != nil {
		fmt.Println(err.Error())
	}


		countryId, err := h.CountryRepository.GetCountry(request.CountryId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
					Status:  http.StatusInternalServerError,
					Message: err.Error(),
				})
				fmt.Println("error 1")
			}	
		trip := models.Trip{
		Title:      	request.Title,
		CountryId:      request.CountryId,
		Country:        ConvertCountryResponse(countryId),
		Accomodation:   request.Accomodation,
		Transportation: request.Transportation,
		Eat:            request.Eat,
		Day:            request.Day,
		Night:          request.Night,
		DateTrip:       request.DateTrip,
		Price:          request.Price,
		Quota:          request.Quota,
		Description:    request.Description,
		Image:          resp.SecureURL,
		}

		data, err := h.TripRepository.CreateTrip(trip)

		if err != nil {
			c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "data udah nambah nyeet", Data: convertResponseTrip(data)})

	} else {
		c.JSON(http.StatusUnauthorized, resultdto.ErrorResult{Status: http.StatusUnauthorized, Message: "Sorry, you're not Admin"})
		return
	}
		
}
func (h *handlerTrip) UpdateTrip(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	UserLogin := c.MustGet("userLogin")
	userAdmin := UserLogin.(jwt.MapClaims)["is_admin"].(bool)
	dataFile := c.MustGet("dataFile").(string)
	fmt.Println("this is data file update", dataFile)

	if userAdmin {
		
			CountryId, _ := strconv.Atoi(c.Request.FormValue("country_id"))
			Day, _ := strconv.Atoi(c.Request.FormValue("day"))
			Night, _ := strconv.Atoi(c.Request.FormValue("night"))
			Price, _ := strconv.Atoi(c.Request.FormValue("price"))
			Quota, _ := strconv.Atoi(c.Request.FormValue("quota"))

		request := tripdto.Update1TripRequest{
			Title: c.Request.FormValue("title"),
			CountryId: CountryId,
			Accomodation: c.Request.FormValue("accomodation"),
			Transportation: c.Request.FormValue("transportation"),
			Eat: c.Request.FormValue("eat"),
			Day: Day,
			Night: Night,
			DateTrip: c.Request.FormValue("date_trip"),
			Price: Price,
			Quota: Quota,
			Description: c.Request.FormValue("description"),
			Image: dataFile,
		}

		if err := c.Bind(request); err != nil {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		trip, err := h.TripRepository.GetTrip(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status:  http.StatusBadRequest, 	Message: err.Error()})
			return
		}

		if request.Title != "" {
			trip.Title = request.Title
		}
		if request.CountryId != 0 {
			trip.CountryId = request.CountryId
		}
		if request.Accomodation != "" {
			trip.Accomodation = request.Accomodation
		}
		if request.Transportation != "" {
			trip.Transportation = request.Transportation
		}
		if request.Eat != "" {
			trip.Eat = request.Eat
		}
		if request.Day != 0 {
			trip.Day = request.Day
		}
		if request.Night != 0 {
			trip.Night = request.Night
		}
		if request.DateTrip != "" {
			trip.DateTrip = request.DateTrip
		}
		if request.Price != 0 {
			trip.Price = request.Price
		}
		if request.Quota != 0 {
			trip.Quota = request.Quota
		}
		if request.Description != "" {
			trip.Description = request.Description
		}
		if request.Image != "" {
			trip.Image = request.Image
		}

		data, err := h.TripRepository.UpdateTrip(trip)

		if err != nil {
			c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "data udeh berhasil lu update nyeeet",Data: convertResponseTrip(data)})
		return

	} else {
		c.JSON(http.StatusUnauthorized, resultdto.ErrorResult{Status: http.StatusUnauthorized, Message: "Sorry, you're not Admin"})
		return
	}
		
}
func (h *handlerTrip) DeleteTrip(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	UserLogin := c.MustGet("userLogin")
	userAdmin := UserLogin.(jwt.MapClaims)["is_admin"].(bool)
	trip, err := h.TripRepository.GetTrip(id)

	if userAdmin {
		if err != nil {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		data, err := h.TripRepository.DeleteTrip(trip)

		if err != nil {
			c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "data udeh berhasil lu apus nyeet", Data: convertResponseTrip(data)})
		return

	} else {
		c.JSON(http.StatusUnauthorized, resultdto.ErrorResult{Status: http.StatusUnauthorized, Message: "Sorry, you're not Admin"})
		return
	}
		
}

func convertResponseTrip(trip models.Trip)tripdto.TripResponse{
	return tripdto.TripResponse{
		ID: trip.ID,
		Title: trip.Title,
		CountryID: trip.CountryId,
		Country: trip.Country,
		Accomodation: trip.Accomodation,
		Transportation: trip.Transportation,
		Eat: trip.Eat,
		Day: trip.Day,
		Night: trip.Night,
		DateTrip: trip.DateTrip,
		Price: trip.Price,
		Quota: trip.Quota,
		Description: trip.Description,
		Image: trip.Image,
	}
}

func convertResponseTrips(trips []models.Trip) []tripdto.TripResponse {
	var responseTrips []tripdto.TripResponse

	for _, trip := range trips {
		responseTrips = append(responseTrips, convertResponseTrip(trip))
	}
	return responseTrips
}

func ConvertCountryResponse(c models.Country) models.CountryResponse {
	return models.CountryResponse{
		ID: c.ID,
		Name: c.Name,
	}
}
