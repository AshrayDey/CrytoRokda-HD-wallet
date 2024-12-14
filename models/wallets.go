package models

import "time"

type Wallet struct {
	ID         int64     `gorm:"primaryKey ; autoIncrement"`
	UserID     int64     `grom:"not null"`
	Mnemonnic  string    `gorm:"not null"`
	MpublicKey string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdateAt   time.Time `gorm:"autoUpdateTime"`
}
