package models

import (
	"github.com/google/uuid"
)

type CinemaChain struct {	
	Name string
}

type CinemaChainDB struct {
	Id uuid.UUID
	CinemaChain *CinemaChain
}