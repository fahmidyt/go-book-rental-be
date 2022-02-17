package models

type Role struct {
	Base
	Name string `gorm:"unique;size:50" json:"name"`
}
