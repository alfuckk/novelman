package authmodel

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	AdminName string         `gorm:"admin_name"`
	Password  string         `gorm:"password"`
	Email     string         `gorm:"email"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Roles     []Role         `gorm:"many2many:admin_roles;"` // 多对多关联到角色表
}

func (Admin) TableName() string {
	return "admins"
}
