package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MagicToken struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID   uuid.UUID `gorm:"type:char(36)"`
	Token    string    `gorm:"uniqueIndex"`
	ExpiresAt time.Time
	Used 	bool `gorm:"default:false"`
	CreatedAt time.Time
}

func (mt *MagicToken) BeforeCreate(tx *gorm.DB) (err error) {
	mt.ID = uuid.New()
	return
}