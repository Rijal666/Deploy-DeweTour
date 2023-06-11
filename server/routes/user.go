package routes

import (
	"dewetour/handlers"
	"dewetour/pkg/middleware"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.GET("/users", middleware.Auth(h.FindUsers))
	r.GET("/user", middleware.Auth(h.GetUser))
	r.PATCH("/user", middleware.Auth(middleware.UploadFile(h.UpdateUser)))
}