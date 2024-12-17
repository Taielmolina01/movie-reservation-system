package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {	
	Email string `gorm:"type:varchar(100);not null"`
	Name string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(100);not null"`
}

type UserDB struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	User User `gorm:"embedded"`
	gorm.Model
}