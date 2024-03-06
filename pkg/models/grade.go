package models

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	StudentID uint
	SubjectID uint
	Grade     float32
}
