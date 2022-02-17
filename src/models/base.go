package models

import (
	"time"

	"gorm.io/gorm"
)

// basicly you can add base columns into all models
type Base struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

// also you can add universal hooks here
// just incase you needed
