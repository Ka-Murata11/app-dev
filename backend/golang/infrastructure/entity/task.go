package entity

import (
	"time"
)

type Task struct {
	TaskID    string    `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
