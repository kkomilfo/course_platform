package models

type Administrator struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
