package data

import "github.com/jinzhu/gorm"

type Reader struct {
	gorm.Model
	FirstName     string `json:"firstname" binding:"required"`
	LastName      string `json:"lastname" binding:"required"`
	Email         string
	UserName      string
	Password      string
	ReaderAccount []Account
}
