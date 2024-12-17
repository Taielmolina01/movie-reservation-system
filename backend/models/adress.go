package models

import (
	"github.com/google/uuid"
)

type Adress struct {
	StreetName string
	HouseNumber int
	City string
	State string
	Country string
	PostalCode int
}

type AdressDB struct {
	Id uuid.UUID
	Adress *Adress
}