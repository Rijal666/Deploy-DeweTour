package handlers

import (
	resultdto "dewetour/dto/result"
	transactiondto "dewetour/dto/transaction"
	"dewetour/models"
	"dewetour/repositories"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"gopkg.in/gomail.v2"
)

type HandleTransaction struct {
	TransactionRepository repositories.TransactionRepository
	TripRepository repositories.TripRepository
	userRepository repositories.UserRepository
}

func NewHandleTransaction(TransactionRepository repositories.TransactionRepository, TripRepository repositories.TripRepository, UserRepository repositories.UserRepository) *HandleTransaction {
	return &HandleTransaction{ TransactionRepository,TripRepository, UserRepository}
}

func (h *HandleTransaction) GetTransaction(c *gin.Context) {
	userLogin := c.MustGet("userLogin")
	userId,_ := userLogin.(jwt.MapClaims)["id"].(float64)

	transaction, err := h.TransactionRepository.GetTransaction(int(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "transaction ", Data: transaction})
}

func (h *HandleTransaction) CreateTransaction(c *gin.Context) {
	// var request transactiondto.TransRequest
	// if err := c.ShouldBindJSON(&request); err != nil {
	// 	c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	// 	return
	userLogin := c.MustGet("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	counterqty, _ := strconv.Atoi(c.PostForm("counter_qty"))
	total, _ := strconv.Atoi(c.PostForm("total"))
	tripid, _ := strconv.Atoi(c.PostForm("trip_id"))

	request := transactiondto.TransactionResponse{
		Name: c.PostForm("name"),
		Gender: c.PostForm("gender"),
		Phone: c.PostForm("phone"),
		CounterQty: counterqty,
		Total:      total,
		Status:     c.PostForm("status"),
		TripID:     tripid,
		UserID:     int(userId),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tripID, err := h.TripRepository.GetTrip(request.TripID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		fmt.Println("error 1")
	}	
	user, _ := h.userRepository.GetUser(request.UserID)

// 	var transactionIsMatch = false
// var transactionId int
// for !transactionIsMatch {
//   transactionId = int(time.Now().Unix())
//   transactionData, _ := h.TransactionRepository.GetTransaction(transactionId)
//   if transactionData.ID == 0 {
//     transactionIsMatch = true
//   }
// }

	transaction := models.Transaction{
		Name: request.Name,
		Gender: request.Gender,
		Phone: request.Phone,
		CounterQty: request.CounterQty,
		Total:      request.Total,
		Status:     request.Status,
		TripId:     request.TripID,
		Trip: CovertTripResponse(tripID),
		UserID:     request.UserID,
		User: ConvertResponseUser(user),
	}

	data, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

// 	// 1. Initiate Snap client
// var s = snap.Client{}
// s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
// // Use to midtrans.Production if you want Production Environment (accept real transaction).

// // 2. Initiate Snap request param
// req := &snap.Request{
//   TransactionDetails: midtrans.TransactionDetails{
//     OrderID:  strconv.Itoa(data.ID),
//     GrossAmt: int64(data.Total),
//   },
//   CreditCard: &snap.CreditCardDetails{
//     Secure: true,
//   },
//   CustomerDetail: &midtrans.CustomerDetails{
//     FName: data.Name,
//     Email: data.Gender,
//     Phone: data.Phone,
//   },
// }

// // 3. Execute request create Snap transaction to Midtrans Snap API
// snapResp, _ := s.CreateTransaction(req)

 c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Data: data})

}

func (h *HandleTransaction) Notification(c *gin.Context) {
	var notificationPayload map[string]interface{}
  
	if err := c.Bind(&notificationPayload); err != nil {
	   c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	   return
	}
  
	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)
  
	order_id, _ := strconv.Atoi(orderId)
  
	fmt.Print("ini payloadnya", notificationPayload)
	transaction, _ := h.TransactionRepository.GetTransaction(order_id)
  
	if transactionStatus == "capture" {
	  if fraudStatus == "challenge" {
		// TODO set transaction status on your database to 'challenge'
		// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
		h.TransactionRepository.UpdateTransaction("pending", order_id)
	  } else if fraudStatus == "accept" {
		// TODO set transaction status on your database to 'success'
		SendMail("success", transaction)
		h.TransactionRepository.UpdateTransaction("success", order_id)
	  }
	} else if transactionStatus == "settlement" {
	  // TODO set transaction status on your databaase to 'success'
	  SendMail("success", transaction)
	  h.TransactionRepository.UpdateTransaction("success", order_id)
	} else if transactionStatus == "deny" {
	  // TODO you can ignore 'deny', because most of the time it allows payment retries
	  // and later can become success
	  h.TransactionRepository.UpdateTransaction("failed", order_id)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
	  // TODO set transaction status on your databaase to 'failure'
	  h.TransactionRepository.UpdateTransaction("failed", order_id)
	} else if transactionStatus == "pending" {
	  // TODO set transaction status on your databaase to 'pending' / waiting payment
	  h.TransactionRepository.UpdateTransaction("pending", order_id)
	}
  
	 c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Data: notificationPayload})
  }

func (h *HandleTransaction) FindTransaction(c *gin.Context) {

	transactions, err := h.TransactionRepository.FindTransaction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	

	c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Data: transactions})
}

func (h *HandleTransaction) DeleteTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	data, err := h.TransactionRepository.DeleteTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Data: data})
}

func SendMail(status string, transaction models.Transaction) {

	if status != transaction.Status && (status == "success") {
	  var CONFIG_SMTP_HOST = "smtp.gmail.com"
	  var CONFIG_SMTP_PORT = 587
	  var CONFIG_SENDER_NAME = "DumbMerch <demo.dumbways@gmail.com>"
	  var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
	  var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")
  
	  var tripName = transaction.Trip.Title
	  var price = strconv.Itoa(transaction.Trip.Price)
  
	  mailer := gomail.NewMessage()
	  mailer.SetHeader("From", CONFIG_SENDER_NAME)
	  mailer.SetHeader("To", transaction.User.Email)
	  mailer.SetHeader("Subject", "Transaction Status")
	  mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
	  <html lang="en">
		<head>
		<meta charset="UTF-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Document</title>
		<style>
		  h1 {
		  color: brown;
		  }
		</style>
		</head>
		<body>
		<h2>Product payment :</h2>
		<ul style="list-style-type:none;">
		  <li>Name : %s</li>
		  <li>Total payment: Rp.%s</li>
		  <li>Status : <b>%s</b></li>
		</ul>
		</body>
	  </html>`, tripName, price, status))
  
	  dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	  )
  
	  err := dialer.DialAndSend(mailer)
	  if err != nil {
		log.Fatal(err.Error())
	  }
  
	  log.Println("Mail sent! to " + transaction.User.Email)
	}
  }

func CovertTripResponse(c models.Trip) models.TripResponse{
	return models.TripResponse{
		ID: c.ID,
		Title: c.Title,
		CountryID: c.CountryId,
		Country: c.Country,
		Accomodation: c.Accomodation,
		Transportation: c.Transportation,
		Eat: c.Eat,
		Day: c.Day,
		Night: c.Night,
		DateTrip: c.DateTrip,
		Price: c.Price,
		Quota: c.Quota,
		Description: c.Description,
		Image: c.Image,
	}
}

func ConvertResponseUser(u models.User) models.UsersProfileResponse  {
	return models.UsersProfileResponse{
		ID: u.ID,
		Fullname: u.Fullname,
		Email: u.Email,
		Phone: u.Phone,
		Address: u.Address,
	}
	
}