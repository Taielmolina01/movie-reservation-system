package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SeatRow int

const (
	A SeatRow = iota
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M
	N
	O
	P
	Q
	R
	S
	T
)

func (s SeatRow) String() string {
	return []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
		"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	}[s]
}

func (s SeatRow) EnumIndex() int {
	return int(s)
}

type SeatRequest struct {
	Row    SeatRow `json:"seatRow" binding:"required"`
	Number int     `json:"number"`
}

type SeatDB struct {
	ID           uuid.UUID    `gorm:"type:uuid;primaryKey"`
	Row          string       `gorm:"type:varchar(2);not null"`
	Number       int          `gorm:"type:int"`
	CinemaRoomID uuid.UUID    `gorm:"type:uuid;not null"`
	CinemaRoom   CinemaRoomDB `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	gorm.Model
}
