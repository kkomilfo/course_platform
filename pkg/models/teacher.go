package models

import (
	"gorm.io/gorm"
	"time"
)

type Teacher struct {
	gorm.Model
	Email       string    `gorm:"unique;not null"`
	Password    string    `gorm:"not null"`
	FullName    string    `gorm:"not null"`
	AvatarURL   string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	DateOfBirth time.Time `gorm:"not null"`
	Education   string    `gorm:"not null"`
}
