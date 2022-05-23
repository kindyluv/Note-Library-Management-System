package data

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	FirstName   string `json:"firstname" binding:"required"`
	LastName    string `json:"lastname" binding:"required"`
	Email       string `gorm:"unique"`
	Password    string
	Age         int64
	Sex         string
	AccountType AccountType
}
