package repositories

import (
	"dewetour/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransaction() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, orderId int) (models.Transaction, error)}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Trip.Country").Preload("User").Find(&transactions).Error

	return transactions, err
}
func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Trip.Country").Preload("User").First(&transaction, ID).Error

	return transaction, err
}
func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Trip").Preload("Coutry").Create(&transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(status string, orderId int) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("User").Preload("Trip.Country").First(&transaction, orderId)
  
	if status != transaction.Status && status == "success" {
	  var transaction models.Transaction
	  r.db.First(&transaction, transaction.Trip.ID)
	  transaction.CounterQty = transaction.CounterQty - 1
	  r.db.Save(&transaction)
	}
  
	transaction.Status = status
	err := r.db.Save(&transaction).Error
	return transaction, err
  }