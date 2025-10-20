package domain

import (
	"gorm.io/gorm"
)

// Entity is the base for all persisted models. It embeds gorm.Model to support
// ID, timestamps and soft delete handling.
type Entity struct {
	gorm.Model
}
