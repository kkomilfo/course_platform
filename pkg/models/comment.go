package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID    uint   `gorm:"not null"`
	UserType  string `gorm:"type:varchar(10);not null"` // 'student' or 'teacher'
	SubjectID uint   `gorm:"not null"`
	Content   string `gorm:"not null"`
	Subject   Subject
}
