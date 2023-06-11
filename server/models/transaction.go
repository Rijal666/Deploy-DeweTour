package models

import "time"

type Transaction struct {
	ID         int            `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name" gorm:"type varchar(255)"`
	Gender string `json:"gender" gorm:"type varchar(255)"`
	Phone string `json:"phone" gorm:"type varchar(255)"`
	CounterQty int            `json:"counter_Qty" gorm:"type int"`
	Total      int            `json:"total" gorm:"type int"`
	Status     string         `json:"status" gorm:"type varchar(255)"`
	TripId     int                `json:"tripid" form:"tripid"`
	Trip       TripResponse       `json:"trip"`
	UserID int `json:"user_id"`
	User UsersProfileResponse `json:"user"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
}

type TransactionResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Phone string `json:"phone"`
	CounterQty int            `json:"counter_Qty"`
	Total      int            `json:"total"`
	Status     string         `json:"status"`
	UserID int `json:"-"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}