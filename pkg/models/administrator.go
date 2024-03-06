package models

import "gorm.io/gorm"

type Administrator struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
