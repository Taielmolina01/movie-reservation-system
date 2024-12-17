package models

import (
	"github.com/google/uuid"
)

type CinemaHeadquarter struct {	
	Name string
}

type CinemaHeadquarterDB struct {
	Id uuid.UUID
	CinemaHeadQuarter *CinemaHeadquarter
	// CinemaChainID
	// HeadquarterAdressID
}