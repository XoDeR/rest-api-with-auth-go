package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string     `json:"name"`
	Email     string     `json:"email" gorm:"unique;not null"`
	Password  string     `json:"-"`
	Tags      []Tag      `gorm:"foreignKey:UserID"`
	TimeSpans []TimeSpan `gorm:"foreignKey:UserID"`
}
