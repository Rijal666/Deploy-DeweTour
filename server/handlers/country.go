package handlers

import (
	countrydto "dewetour/dto/country"
	resultdto "dewetour/dto/result"
	"dewetour/models"
	"dewetour/repositories"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerCountry struct {
	CountryRepository repositories.CountryRepository
}

func HandlerCountry(CountryRepository repositories.CountryRepository) *handlerCountry {
	return &handlerCountry{CountryRepository}
}

func (h *handlerCountry) FindCountries(c *gin.Context){
	userLogin := c.MustGet("userLogin")
	userAdmin := userLogin.(jwt.MapClaims)["is_admin"].(bool)

	if userAdmin {
		countries, err := h.CountryRepository.FindCountries()
		if  err != nil {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}
		if len(countries) > 0 {
			c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Semua data berhasil ditampilan, ok nyeet", Data: countries})
		} else {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: "Data kosong, tambah dulu nyeet"})
		}

	} else {
		c.JSON(http.StatusUnauthorized, resultdto.ErrorResult{Status: http.StatusUnauthorized, Message: "Lu bukan admin nyeet"})
		return
	}
	
}
func (h *handlerCountry) GetCountry(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
fmt.Println(id)
	
	country, err := h.CountryRepository.GetCountry(id)
	fmt.Println(country)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Country data successfully obtained", Data: country})

		} 

func (h *handlerCountry) CreateCountry(c *gin.Context){
	userLogin := c.MustGet("userLogin")
	userAdmin := userLogin.(jwt.MapClaims)["is_admin"].(bool)

	if userAdmin {
		request := new(countrydto.CreateCountryRequest)
		if err := c.Bind(request); err != nil {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		validation := validator.New()
		err := validation.Struct(request)

		if err != nil {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		country := models.Country{
			Name: request.Name,
		}

		data, err := h.CountryRepository.CreateCountry(country)

		if err != nil {
			c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "data udah nambah nyeet",Data: convertResponseCountry(data)})
		return

		} else {
			c.JSON(http.StatusUnauthorized, resultdto.ErrorResult{Status: http.StatusUnauthorized, Message: "Lu bukan admin nyeet"})
			return
		}
}

func (h *handlerCountry) UpdateCountry(c *gin.Context){
	userLogin := c.MustGet("userLogin")
	userAdmin := userLogin.(jwt.MapClaims)["is_admin"].(bool)
	
	if userAdmin {
		request := new(countrydto.UpdateCountryRequset)

		if err := c.Bind(request); err != nil {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		id, _ := strconv.Atoi(c.Param("id"))
		country, err := h.CountryRepository.GetCountry(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		if request.Name != "" {
			country.Name = request.Name
		}

		country.UpdatedAt = time.Now()

		data, err := h.CountryRepository.UpdateCountry(country)

		if err != nil {
			c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "data udeh berhasil lu update nyeeet",Data: convertResponseCountry(data)})
		return

		} else {
			c.JSON(http.StatusUnauthorized, resultdto.ErrorResult{Status: http.StatusUnauthorized, Message: "Lu bukan admin nyeet"})
			return
		}
}
func (h *handlerCountry) DeleteCountry(c *gin.Context){
	userLogin := c.MustGet("userLogin")
	userAdmin := userLogin.(jwt.MapClaims)["is_admin"].(bool)
	id, _ := strconv.Atoi(c.Param("id"))
	country, err := h.CountryRepository.GetCountry(id)

	if userAdmin {
		if err != nil {
			 c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
				Status:    http.StatusBadRequest,
				Message: err.Error()})
				return
		}

		data, err := h.CountryRepository.DeleteCountry(country)

		if err != nil {
			c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "data udeh berhasil lu apus nyet",Data: convertResponseCountry(data)})
		return

		} else {
			c.JSON(http.StatusUnauthorized, resultdto.ErrorResult{Status: http.StatusUnauthorized, Message: "Lu bukan admin nyeet"})
			return
		}
}

func convertResponseCountry(country models.Country) countrydto.CountryResponse{
	return countrydto.CountryResponse{
		ID: country.ID,
		Name: country.Name,
	}
}