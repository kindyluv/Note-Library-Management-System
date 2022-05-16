package data

import "github.com/jinzhu/gorm"

type RolePermission struct {
	gorm.Model
	RoleID       uint
	PermissionID uint
}

func (r RolePermission) TableName() string {
	return tablePrefix + "role_permissions"
}
