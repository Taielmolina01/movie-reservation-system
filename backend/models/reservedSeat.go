package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservedSeatDB struct {
	ID            uuid.UUID     `gorm:"type:uuid;primaryKey"`
	SeatID        uuid.UUID     `gorm:"type:uuid;not null"`
	Seat          SeatDB        `gorm:"foreignKey:SeatID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ReservationID uuid.UUID     `gorm:"type:uuid;not null"`
	Reservation   ReservationDB `gorm:"foreignKey:ReservationID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	gorm.Model
}
