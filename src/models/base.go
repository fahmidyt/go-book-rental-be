package models

import (
	"gorm.io/gorm"
)

// basicly you can add base columns into all models
type Base struct {
	gorm.Model
}

// also you can add universal hooks here
// just incase you needed
