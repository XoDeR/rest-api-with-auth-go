package models

import (
	"time"

	"gorm.io/gorm"
)

type TimeSpan struct {
	gorm.Model
	Start  time.Time
	End    time.Time
	UserID uint  `gorm:"not null"`
	User   User  `gorm:"constraint:OnDelete:CASCADE;"`
	Tags   []Tag `gorm:"many2many:timespan_tags"`
}
