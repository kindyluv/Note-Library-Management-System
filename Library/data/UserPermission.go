package data

import "github.com/jinzhu/gorm"

type UserPermission struct {
	gorm.Model
	Name string
}

func (permission UserPermission) TableName() string {
	return tablePrefix + "permissions"
}
