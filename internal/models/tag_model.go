package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Color  string `gorm:"not null" json:"color"`
	Key    string `gorm:"not null" json:"key"`
	UserID uint   `gorm:"not null"`
	User   User   `gorm:"constraint:OnDelete:CASCADE;"`
}
