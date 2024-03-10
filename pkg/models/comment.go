package models

type Comment struct {
	ID            uint   `gorm:"primarykey" json:"id"`
	UserID        uint   `gorm:"not null" json:"user_id"`
	UserType      string `gorm:"type:varchar(10);not null" json:"user_type"`
	StudentWorkID uint   `gorm:"not null" json:"student_work_id"`
	Content       string `gorm:"not null" json:"content"`
}

type StudentWork struct {
	ID        uint              `gorm:"primarykey" json:"id"`
	SubjectID uint              `gorm:"not null" json:"subject_id"`
	StudentID uint              `gorm:"not null" json:"student_id"`
	Grade     *uint             `json:"grade"`
	Comments  []Comment         `gorm:"foreignKey:StudentWorkID" json:"comments"`
	Files     []StudentWorkFile `gorm:"foreignKey:StudentWorkID" json:"files"`
}

type StudentWorkFile struct {
	ID            uint   `gorm:"primarykey" json:"id"`
	StudentWorkID uint   `gorm:"not null" json:"student_work_id"`
	Name          string `json:"name"`
	URL           string `gorm:"not null" json:"url"`
}
