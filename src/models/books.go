package models

import "time"

type Book struct {
	Base
	Name         string    `gorm:"not null;size:100" json:"name"`
	Author       string    `gorm:"not null;size:100" json:"author"`
	ReleasedDate time.Time `json:"releasedDate"`
}
