package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CinemaChain struct {
	Name string `gorm:"type:varchar(100);not null"`
}

type CinemaChainDB struct {
	ID          uuid.UUID   `gorm:"type:uuid;primaryKey"`
	CinemaChain CinemaChain `gorm:"embedded"`
	gorm.Model
}
