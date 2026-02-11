package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name string    `json:"name" gorm:"unique;not null"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}
