package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservedSeatRequest struct {
	ID            uuid.UUID `json:"number"`
	SeatID        uuid.UUID `json:"number"`
	ReservationID uuid.UUID `json:"number"`
}

type ReservedSeatDB struct {
	ID            uuid.UUID     `gorm:"type:uuid;primaryKey"`
	SeatID        uuid.UUID     `gorm:"type:uuid;not null"`
	Seat          SeatDB        `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ReservationID uuid.UUID     `gorm:"type:uuid;not null"`
	Reservation   ReservationDB `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	gorm.Model
}
