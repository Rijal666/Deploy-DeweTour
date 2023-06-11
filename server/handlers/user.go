package handlers

import (
	resultdto "dewetour/dto/result"
	userdto "dewetour/dto/user"
	"dewetour/models"
	"dewetour/repositories"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type handler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{
		UserRepository:        UserRepository,
		// CartRepository:        CartRepository,
		// TransactionRepository: TransactionRepository,
	}
}

func (h *handler) FindUsers(c *gin.Context) {
	userLogin := c.MustGet("userLogin")
	userAdmin := userLogin.(jwt.MapClaims)["is_admin"].(bool)
	if userAdmin {
		users, err := h.UserRepository.FindUsers()
		if err != nil {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		if len(users) > 0 {
			c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Data for all users was successfully obtained", Data: users})
			return
		} else {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: "No record found"})
			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, resultdto.ErrorResult{Status: http.StatusUnauthorized, Message: "Sorry, you're not Admin"})
	}
}

func (h *handler) GetUser(c *gin.Context) {
	userLogin := c.MustGet("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, err := h.UserRepository.GetUser(int(userId))
	if err != nil {
		 c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	 c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "User data successfully obtained", Data: user})
}

func (h *handler) UpdateUser(c *gin.Context){
	dataFile := c.MustGet("dataFile").(string)
		fmt.Println("this is data file", dataFile)
	userLogin := c.MustGet("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)
	
		request := userdto.UpdateUserRequest{
			Fullname: c.Request.FormValue("fullname"),
			Email: c.Request.FormValue("email"),
			Password: c.Request.FormValue("password"),
			Phone: c.Request.FormValue("phone"),
			Address: c.Request.FormValue("address"),
			Image: dataFile,
		}
		
		if err := c.Bind(request); err != nil {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}
		
		user, err := h.UserRepository.GetUser(int(userId))
		fmt.Println(user, "inini tood")

		if err != nil {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		if request.Fullname != "" {
			user.Fullname = request.Fullname
		}
		if request.Email != "" {
			user.Email = request.Email
		}
		if request.Password != "" {
			user.Password = request.Password
		}
		if request.Phone != "" {
			user.Phone = request.Phone
		}
		if request.Address != "" {
			user.Address = request.Address
		}
		if request.Image != "" {
			user.Image = request.Image
		}

		data, err := h.UserRepository.UpdateUser(user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "data udeh berhasil lu update nyeeet",Data: convertResponse(data)})
}


func  convertResponse(u models.User)  userdto.UserResponse{
	return userdto.UserResponse{
		ID: u.ID,
		Fullname: u.Fullname,
		Email: u.Email,
		Phone: u.Phone,
		Address: u.Address,
		Image: u.Image,
	}
	
}