package models

import (
	"time"
)

type UserDetail struct {
	Base
	UserID      uint
	firstName   string    `gorm:"not null"`
	lastName    string    `gorm:"not null"`
	phoneNumber string    `gorm:"not null"`
	dateOfBirth time.Time `gorm:"not null"`
	User        User
}
