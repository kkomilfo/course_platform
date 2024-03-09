package models

type Comment struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	UserID    uint   `gorm:"not null" json:"user_id"`
	UserType  string `gorm:"type:varchar(10);not null" json:"user_type"`
	SubjectID uint   `gorm:"not null" json:"subject_id"`
	Content   string `gorm:"not null" json:"content"`
	Subject   Subject
}
