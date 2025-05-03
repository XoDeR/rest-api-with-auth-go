package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Color string `gorm:"not null" json:"color"`
	Key   string `gorm:"not null" json:"key"`
	User  User   `gorm:"foreignkey:UserID"`
}
