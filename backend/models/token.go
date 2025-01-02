package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type TokenDB struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	AccessToken  string    `gorm:"type:varchar(255);not null"`
	RefreshToken string    `gorm:"type:varchar(255);not null"`
	UserEmail    string    `gorm:"type:varchar(255);not null"`
	ExpiresAt    time.Time `gorm:"not null"`
	User         UserDB    `gorm:"foreignKey:Email;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	gorm.Model
}
