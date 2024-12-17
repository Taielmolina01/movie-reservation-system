package models 

func GetAllModels() []interface{} {
	return []interface{}{
		&AdressDB{},
		&CinemaChainDB{},
		&CinemaHeadquarterDB{},
		&CinemaRoomDB{},
		&CinemaShowDB{},
		&MovieDB{},
		&ReservationDB{},
		&ReservedSeatDB{},
		&SeatDB{},
		&TokenDB{},
		&UserDB{},
	}
}