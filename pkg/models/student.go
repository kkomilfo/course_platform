package models

type Student struct {
	ID        uint     `gorm:"primarykey" json:"id"`
	Email     string   `gorm:"unique;not null" json:"email"`
	Password  string   `gorm:"not null" json:"-"`
	AvatarURL string   `gorm:"not null" json:"avatar_url"`
	FullName  string   `gorm:"not null" json:"full_name"`
	Courses   []Course `gorm:"many2many:course_enrollments;" json:"courses"`
}
