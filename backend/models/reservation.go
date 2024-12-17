package models

import (
	"github.com/google/uuid"
	"time"
)

type Reservation struct {	
	Date time.Time
}

type ReservationDB struct {
	Id uuid.UUID
	Reservation *Reservation
	// UserId
	// CinemaShowID
}