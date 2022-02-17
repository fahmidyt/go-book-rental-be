package models

import (
	"time"
)

type UserDetail struct {
	Base
	UserID      uint
	FirstName   string    `gorm:"not null;size:100" json:"firstName"`
	LastName    string    `gorm:"not null;size:100" json:"lastName"`
	PhoneNumber string    `gorm:"not null;size:15" json:"phoneNumber"`
	DateOfBirth time.Time `gorm:"not null" json:"dateOfBirth"`
}
