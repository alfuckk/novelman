package authmodel

import "gorm.io/gorm"

type RolePermission struct {
	gorm.Model
	RoleID       uint
	PermissionID uint
	Role         Role
	Permission   Permission
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
