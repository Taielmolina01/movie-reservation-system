package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MovieGenre int

const (
	Action MovieGenre = iota
	Adventure
	Animation
	Biography
	Comedy
	Crime
	Documentary
	Drama
	Fantasy
	Family
	History
	Horror
	Musical
	Mystery
	Romance
	ScienceFiction
	Sports
	Thriller
	War
	Western
)

func (mg MovieGenre) String() string {
	return [...]string{
		"Action",
		"Adventure",
		"Animation",
		"Biography",
		"Comedy",
		"Crime",
		"Documentary",
		"Drama",
		"Fantasy",
		"Family",
		"History",
		"Horror",
		"Musical",
		"Mystery",
		"Romance",
		"Science Fiction",
		"Sports",
		"Thriller",
		"War",
		"Western",
	}[mg]
}

func (mg MovieGenre) EnumIndex() int {
	return int(mg)
}

type Movie struct {
	Title       string     `gorm:"type:varchar(100);not null"`
	Description string     `gorm:"type:varchar(255);not null"`
	Poster      []byte     `gorm:"type:bytea;not null"`
	Genre       MovieGenre `gorm:"type:varchar(50);not null`
}

type MovieDB struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Movie Movie     `gorm:"embedded"`
	gorm.Model
}
