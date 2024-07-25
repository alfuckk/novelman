package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
}

func (m *Admin) TableName() string {
    return "admin"
}
