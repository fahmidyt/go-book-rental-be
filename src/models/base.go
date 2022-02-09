package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID `gorm:"uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(db *gorm.DB) error {
	uuid := uuid.New().String()

	db.Statement.SetColumn("ID", uuid)
	return nil
}
