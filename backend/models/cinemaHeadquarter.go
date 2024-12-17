package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CinemaHeadquarter struct {	
	Name string `gorm:"type:varchar(100);not null"`
}

type CinemaHeadquarterDB struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	CinemaHeadquarter CinemaHeadquarter `gorm:"embedded"`
	// FK 
	CinemaChainID uuid.UUID `gorm:"type:uuid;not null"`
	// Declare the relationship CinemaHeadquarter - CinemaChain.
	// If the CinemaChainID is updated, it will be updated in this table. 
	// By the other hand if I want to delete a CinemaChain having CinemaHeadquarters, i won't be able to do it.
	CinemaChain CinemaChainDB `gorm:"foreignKey:CinemaChainID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	// FK
	HeadquarterAdressID uuid.UUID `gorm:"type:uuid;not null`
	// Declare the relationship CinemaHeadquarter - Adress
	HeadquarterAdress AdressDB `gorm:"foreignKey:HeadquarterAdressID;constraint:OnUpdate:Cascade,OnDelete:RESTRICT"`
	
	gorm.Model
}