package routes

import (
	"dewetour/handlers"
	"dewetour/pkg/middleware"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(r *gin.RouterGroup) {
	TransactionRepository := repositories.RepositoryTransaction(mysql.DB)
	TripRepository := repositories.RepositoryTrip(mysql.DB)
	UserRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.NewHandleTransaction(TransactionRepository, TripRepository, UserRepository)

	r.GET("/transactions", middleware.Auth(h.FindTransaction))
	r.GET("/transaction", middleware.Auth(h.GetTransaction))
	r.POST("/transaction", middleware.Auth(h.CreateTransaction))
	r.DELETE("/transaction/:id", middleware.Auth(h.DeleteTransaction))
	r.POST("/notification", h.Notification)
}