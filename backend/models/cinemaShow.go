package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type CinemaShow struct {
	ShowDate time.Time `gorm:"not null`
}

type CinemaShowDB struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey"`
	CinemaShow CinemaShow `gorm:"embedded"`
	// FK
	MovieID uuid.UUID `gorm:"type:uuid;not null"`
	// Declare relationship CinemaShow - Movie
	Movie MovieDB `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	// FK
	CinemaRoomID uuid.UUID `gorm:"type:uuid;not null"`
	// Declare relationship CinemaShow - CinemaRoom
	CinemaRoom CinemaRoomDB `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	gorm.Model
}
