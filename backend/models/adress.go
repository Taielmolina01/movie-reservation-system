package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdressRequest struct {
	StreetName  string `json:"streetName" binding:"required"`
	HouseNumber int    `json:"houseNumber" binding:"required"`
	City        string `json:"city" binding:"required"`
	State       string `json:"state" binding:"required"`
	Country     string `json:"country" binding:"required"`
	PostalCode  int    `json:"postalCode" binding:"required"`
}

type AdressDB struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	StreetName  string    `gorm:"type:varchar(255);not null"`
	HouseNumber int       `grom:"not null"`
	City        string    `gorm:"type:varchar(100);not null"`
	State       string    `gorm:"type:varchar(100);not null"`
	Country     string    `gorm:"type:varchar(100);not null"`
	PostalCode  int       `gorm:"not null"`
	gorm.Model
}
