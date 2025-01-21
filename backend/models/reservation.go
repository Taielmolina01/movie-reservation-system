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
	UserEmail    string       `gorm:"type:varchar(255);not null"`
	User         UserDB       `gorm:"foreignKey:Email;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CinemaShowID uuid.UUID    `gorm:"type:uuid;not null"`
	CinemaShow   CinemaShowDB `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	gorm.Model
}
