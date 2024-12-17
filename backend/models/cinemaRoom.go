package models

import (
	"github.com/google/uuid"
)

// aca tengo que ver que onda porque perse el headquarter no tendria nada

type CinemaRoom struct {	
	DiagramRoom string
}

type CinemaRoomDB struct {
	Id uuid.UUID
	CinemaRoom *CinemaRoom
	// HeadquarterID
}