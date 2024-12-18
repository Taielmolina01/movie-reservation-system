package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Token struct {
	AccessToken  string    `gorm:"type:varchar(255);not null"`
	RefreshToken string    `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `gorm:"not null"`
	ExpiresAt    time.Time `gorm:"not null"`
}

type TokenDB struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	Token  Token     `gorm:"embedded"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User   UserDB    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	gorm.Model
}
