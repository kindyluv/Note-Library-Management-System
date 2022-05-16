package data

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Title   string `json:"book" binding:"required"`
	ISBN    string `json:"isbn" binding:"required"`
	BookUrl string `json:"book_url" binding:"required"`
	Author  Author
	Edition string `json:"edition" binding:"required"`
}
