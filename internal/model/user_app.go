package model

import "gorm.io/gorm"

// user app
type UserApp struct {
	gorm.Model
	UserID uint
	AppID  uint
	Mark   string `gorm:"unique;not null"`
	Appid  string `gorm:"unique;not null"`
	Appkey string `gorm:"unique;not null"`
	Status int64
	Quota  int64
	User   User
	App    App
}

func (u *UserApp) TableName() string {
	return "user_apps"
}
