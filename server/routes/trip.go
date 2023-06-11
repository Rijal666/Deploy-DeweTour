package routes

import (
	"dewetour/handlers"
	"dewetour/pkg/middleware"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gin-gonic/gin"
)

func TripRoutes(r *gin.RouterGroup) {
	TripRepository := repositories.RepositoryTrip(mysql.DB)
	CountryRepository := repositories.RepositoryCountry(mysql.DB)
	h := handlers.HandlerTrip(TripRepository, CountryRepository)

	r.GET("/trips", h.FindTrips)
	r.GET("/trip/:id", h.GetTrip)
	r.POST("/trip", middleware.Auth(middleware.UploadFile(h.CreateTrip)))
	r.PATCH("/trip/:id", middleware.Auth(middleware.UploadFile(h.UpdateTrip)))
	r.DELETE("/trip/:id", middleware.Auth(h.DeleteTrip))
	
}