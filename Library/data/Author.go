package data

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Location  string
	Book      []Book
}
