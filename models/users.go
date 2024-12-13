package models

type User struct {
	Id       int64  `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"size:255;not null" binding:"required"`
	Email    string `gorm:"size:255;unique;not null" binding:"required,email"`
	Password string `gorm:"size:255;not null" binding:"required"`
}
