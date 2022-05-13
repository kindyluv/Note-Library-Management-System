package data

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Name   string `json:"book" binding:"required"`
	Author []Author
}
