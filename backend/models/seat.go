package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type SeatRow int

const (
	A		SeatRow = iota
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

type Seat struct {	
	Row SeatRow `gorm:"type:varchar(2);not null"`
	Number int `gorm:"not null"`
}

type SeatDB struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Seat Seat `gorm:"embedded"`
	CinemaRoomID uuid.UUID `gorm:"type:uuid;not null"`
	CinemaRoom CinemaRoomDB `gorm:"foreignKey:CinemaRoomID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	gorm.Model
}