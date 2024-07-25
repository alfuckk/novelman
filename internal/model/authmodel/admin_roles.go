package authmodel

import "gorm.io/gorm"

type AdminRole struct {
	gorm.Model
	AdminID uint
	RoleID  uint
	Admin   Admin
	Role    Role
}

func (AdminRole) TableName() string {
	return "admin_roles"
}
