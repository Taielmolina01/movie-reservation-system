package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CinemaRoom struct {
	DiagramRoom []byte `gorm:"type:bytea;not null"`
}

type CinemaRoomDB struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey"`
	CinemaRoom CinemaRoom `gorm:"embedded"`
	// FK
	HeadquarterID uuid.UUID `gorm:"type:uuid;not null"`
	// Declare relationship CinemaRoom - CinemaHeadquarter
	Headquarter CinemaHeadquarterDB `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	gorm.Model
}
