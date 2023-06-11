package migration

import (
	"dewetour/models"
	"dewetour/pkg/mysql"
	"fmt"
)

func RunAutoMigrate() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Country{},
		&models.Trip{},
		&models.Transaction{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Migration")
}