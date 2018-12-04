package room

type Room struct {
	RoomID string
	SeatNumber int
	CreateTime int
	StartTime int
	EndTime int
	//todo
}

type Seat struct {
	RoomID string
	//RC example A1, AA1
	SeatID string
	//available, occupied, hidden, booked
	SeatStatus string
	UserID string
	CreateTime int
}
