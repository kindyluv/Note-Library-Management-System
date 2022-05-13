package data

import "github.com/jinzhu/gorm"

type Reader struct {
	gorm.Model
	UserName      string `gorm:"unique"`
	Password      string
	ReaderAccount []Account
}
