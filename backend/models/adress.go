package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Adress struct {
	StreetName string `gorm:"type:varchar(255);not null"`
	HouseNumber int `grom:"not null"`
	City string	`gorm:"type:varchar(100);not null"`
	State string `gorm:"type:varchar(100);not null)"`
	Country string `gorm:"type:varchar(100);not null"`
	PostalCode int `gorm:"not null"`
}

type AdressDB struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Adress Adress `gorm:"embedded"`
	gorm.Model
} 