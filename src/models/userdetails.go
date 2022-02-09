package models

import (
	"time"
)

type UserDetail struct {
	Base
	UserID      uint
	FirstName   string    `gorm:"not null"`
	LastName    string    `gorm:"not null"`
	PhoneNumber string    `gorm:"not null"`
	DateOfBirth time.Time `gorm:"not null"`
	User        User
}
