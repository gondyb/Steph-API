package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base represents the gorm base model. This model has an UUID as an ID and basic timestamps attributes.
type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate handles the uuid generation when a new model is created.
func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	base.ID = uuid
	return
}
