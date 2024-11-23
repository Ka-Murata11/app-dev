package entity

type User struct {
	UserID   string `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string
	Job      string
}
