package authmodel

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName    string       `gorm:"role_name"`
	Admin       []Admin      `gorm:"many2many:admin_roles;"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

func (Role) TableName() string {
	return "roles"
}
