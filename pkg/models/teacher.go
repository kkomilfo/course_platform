package models

import (
	"time"
)

type Teacher struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Email       string    `gorm:"unique;not null" json:"email"`
	Password    string    `gorm:"not null" json:"password"`
	FullName    string    `gorm:"not null" json:"full_name"`
	AvatarURL   string    `gorm:"not null" json:"avatar_url"`
	Description string    `gorm:"not null" json:"description"`
	DateOfBirth time.Time `gorm:"not null" json:"date_of_birth"`
	Education   string    `gorm:"not null" json:"education"`
	Courses     []Course  `gorm:"foreignKey:TeacherID" json:"courses"`
}
