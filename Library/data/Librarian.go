package data

import "github.com/jinzhu/gorm"

type Librarian struct {
	gorm.Model
	UserName         string `gorm:"unique"`
	Password         string `json:"password" binding:"required"`
	LibrarianAccount []Account
}
