package routes

import (
	"dewetour/handlers"
	"dewetour/pkg/middleware"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	authRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepository)

	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.GET("/check-auth", middleware.Auth(h.CheckAuth))
}