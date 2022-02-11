package models

import "time"

type Book struct {
	Base
	Name         string `gorm:"not null;size:100"`
	Author       string `gorm:"not null;size:100"`
	ReleasedDate time.Time
}
