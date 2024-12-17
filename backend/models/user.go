package models

import (
	"github.com/google/uuid"
)

type User struct {	
	Email string
	Name string
	Password string
}

type UserDB struct {
	Id uuid.UUID
	User *User
}