package models

import "time"

type Order struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	UserID      int       `gorm:"not null"`
	OrderDate   time.Time `gorm:"not null"`
	TotalAmount float64   `gorm:"not null"`
}
