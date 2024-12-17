package models

import (
	"github.com/google/uuid"
)


type MovieGenre int

const (
	Action		MovieGenre = iota
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
	Science Fiction
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
		"Horror"
		"Musical",
		"Mystery",
		"Romance",
		"Science Fiction",
		"Sports",
		"Thriller",
		"War",
		"Western"
		}[mg]
}

func (mg MovieGenre) EnumIndex() int {
	return int(mg)
}

type Movie struct {	
	Title string
	Description string
	Poster string
	Genre MovieGenre
}

type MovieDB struct {
	Id uuid.UUID
	Movie *Movie
}