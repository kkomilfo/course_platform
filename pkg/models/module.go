package models

import "gorm.io/gorm"

type Module struct {
	gorm.Model
	Title    string
	CourseID uint
	Subjects []Subject `gorm:"foreignKey:ModuleID"`
}
