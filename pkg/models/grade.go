package models

type Grade struct {
	ID        uint    `gorm:"primarykey" json:"id"`
	StudentID uint    `json:"student_id"`
	SubjectID uint    `json:"subject_id"`
	Grade     float32 `json:"grade"`
}
