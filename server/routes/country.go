package routes

import (
	"dewetour/handlers"
	"dewetour/pkg/middleware"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gin-gonic/gin"
)

func CountryRoutes(r *gin.RouterGroup) {
	CountryRepository := repositories.RepositoryCountry(mysql.DB)
	h := handlers.HandlerCountry(CountryRepository)

	r.GET("/countries", middleware.Auth(h.FindCountries))
	r.GET("/country/:id", h.GetCountry)
	r.POST("/country", middleware.Auth(h.CreateCountry))
	r.PATCH("/country/:id", middleware.Auth(h.UpdateCountry))
	r.DELETE("/country/:id", middleware.Auth(h.DeleteCountry))

}