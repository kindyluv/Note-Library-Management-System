package dto

import (
	"gorm.io/gorm"
)

type ReaderResponse struct {
	gorm.Model
	UserName string
}
