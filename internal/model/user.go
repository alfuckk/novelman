package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   string `gorm:"unique;not null"`
	Nickname string `gorm:"not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Status   int64
	Quota    int64
	Apps     []*App `gorm:"many2many:user_apps;"` // 多对多关联到角色表
}

func (u *User) TableName() string {
	return "users"
}
