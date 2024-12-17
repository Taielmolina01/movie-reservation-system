package models

import (
	"github.com/google/uuid"
	"time"
)

type CinemaShow struct {	
	ShowDate time.Time
}

type CinemaShowDB struct {
	// MovieID
	// CinemaRoomID
	Id uuid.UUID
	CinemaShow *CinemaShow
}