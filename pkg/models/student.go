package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	AvatarURL string
	FullName  string   `gorm:"not null"`
	Courses   []Course `gorm:"many2many:course_enrollments;"`
}
