package models

import (
	"time"
)

type UserDetail struct {
	Base
	UserID      uint
	FirstName   string    `gorm:"not null;size:100"`
	LastName    string    `gorm:"not null;size:100"`
	PhoneNumber string    `gorm:"not null;size:15"`
	DateOfBirth time.Time `gorm:"not null"`
}
