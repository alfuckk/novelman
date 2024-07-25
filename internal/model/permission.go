package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	PermissionName string `gorm:"permission_name"`
	Description    string `gorm:"permission_description"`
	Roles          []Role `gorm:"many2many:role_permissions;"`
}

func (Permission) TableName() string {
	return "permissions"
}
