package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CinemaHeadquarter struct {
	Name string `json:"name" binding:"required"`
}

type CinemaHeadquarterDB struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string    `gorm:"type:varchar(100);not null"`
	// FK
	CinemaChainID uuid.UUID `gorm:"type:uuid;not null"`
	// Declare the relationship CinemaHeadquarter - CinemaChain.
	// If the CinemaChainID is updated, it will be updated in this table.
	// By the other hand if I want to delete a CinemaChain having CinemaHeadquarters, i won't be able to do it.
	CinemaChain CinemaChainDB `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	// FK
	HeadquarterAdressID uuid.UUID `gorm:"type:uuid;not null`
	// Declare the relationship CinemaHeadquarter - Adress
	HeadquarterAdress AdressDB `gorm:"foreignKey:ID;constraint:OnUpdate:Cascade,OnDelete:RESTRICT"`

	gorm.Model
}
