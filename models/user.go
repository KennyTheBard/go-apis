package models

type User struct {
	ID    int    `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null"`
}
