package models

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
	return [...]string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
		"J",
		"K",
		"L",
		"M",
		"N",
		"O",
		"P",
		"Q",
		"R",
		"S",
		"T"
		}[s]
}

func (s SeatRow) EnumIndex() int {
	return int(s)
}

type Seat struct {	
	Row SeatRow
	Number int
}

type SeatDB struct {
	Id uuid.UUID
	Seat *Seat
	// cinemaRoomID
}