package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	TeacherID   uint      `gorm:"not null"`
	Teacher     Teacher   `gorm:"not null"`
	Modules     []Module  `gorm:"foreignKey:CourseID"`
	Students    []Student `gorm:"many2many:course_enrollments;"`
}
