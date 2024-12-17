package models

import (
	"github.com/google/uuid"
	"time"
)

type Token struct {	
	AccessToken	string
	RefreshToken string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type TokenDB struct {
	Id uuid.UUID
	Token *Token
	// userID
}