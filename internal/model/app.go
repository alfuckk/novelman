package model

import "gorm.io/gorm"

// app
// like sms,mms,email,short_url
type App struct {
	gorm.Model
	AppName string `gorm:"app_name"`
	Status  int64  `gorm:"status"`
}

func (u *App) TableName() string {
	return "apps"
}
