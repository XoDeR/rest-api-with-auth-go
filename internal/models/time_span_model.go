package models

import (
	"time"

	"gorm.io/gorm"
)

type TimeSpan struct {
	gorm.Model
	Start time.Time
	End   time.Time
	User  User `gorm:"foreignkey:UserID"`
	Tags  []Tag
}
