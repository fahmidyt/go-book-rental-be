package models

import "time"

type Book struct {
	Base
	Name         string `gorm:"not null"`
	Author       string `gorm:"not null"`
	ReleasedDate time.Time
}
