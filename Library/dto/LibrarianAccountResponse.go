package dto

import "gorm.io/gorm"

type LibrarianAccountResponse struct {
	gorm.Model
	UserName string
}
