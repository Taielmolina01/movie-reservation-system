package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Reservation struct {
	Date time.Time `gorm:"not null"`
}

type ReservationDB struct {
	ID           uuid.UUID    `gorm:"type:uuid;primaryKey"`
	Reservation  Reservation  `gorm:"embedded"`
	UserID       uuid.UUID    `gorm:"type:uuid;not null"`
	User         UserDB       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CinemaShowID uuid.UUID    `gorm:"type:uuid;not null"`
	CinemaShow   CinemaShowDB `gorm:"foreignKey:CinemaShowID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	gorm.Model
}
