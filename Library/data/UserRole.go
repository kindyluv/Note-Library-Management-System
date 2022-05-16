package data

import "github.com/jinzhu/gorm"

type UserRole struct {
	gorm.Model
	UserID uint
	RoleID uint
}

func (r UserRole) TableName() string {
	return tablePrefix + "user_roles"
}
