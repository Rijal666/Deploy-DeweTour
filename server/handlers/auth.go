package handlers

import (
	authdto "dewetour/dto/auth"
	resultdto "dewetour/dto/result"
	"dewetour/models"
	"dewetour/pkg/bcrypt"
	jwtToken "dewetour/pkg/jwt"
	"dewetour/repositories"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(c *gin.Context) {
	// c.Header("content-Type", "application/json")

	request := new(authdto.AuthRequset)
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

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	user := models.User{
		IsAdmin: request.IsAdmin,
		Fullname: request.Fullname,
		Email: request.Email,
		Password: password,
		Phone: request.Phone,
		Address: request.Address,
	}

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Your registration is successful", Data: data})
	
	
}

func(h *handlerAuth) Login(c *gin.Context) {
	request := new(authdto.LoginRequest)
	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	user := models.User{
		IsAdmin: request.IsAdmin,
		Email: request.Email,
		Password: request.Password,
	}
	user, err := h.AuthRepository.Login(user.Email)
	if err != nil{
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: "wrong email or password"})
		return
	}
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["is_admin"] = user.IsAdmin
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil{
		log.Println(errGenerateToken)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error" : "Unauthorized"})
		return
	}

	loginResponse := authdto.LoginResponse{
		ID: user.ID,
		IsAdmin: user.IsAdmin,
		Email: user.Email,
		Token: token,
	}
	c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "You have successfully logged in", Data: loginResponse})
	
	
}

func (h *handlerAuth) CheckAuth(c *gin.Context) {
	userLogin := c.MustGet("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, _ := h.AuthRepository.CheckAuth(int(userId))
	c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Data: user})
	
}