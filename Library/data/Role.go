package data

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name string
}

func (r Role) TableName() string {
	return tablePrefix + "roles"
}
