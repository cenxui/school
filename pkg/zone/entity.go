package zone

type Zone struct {
	ZoneID string
	PositionNumber int
	CreateTime int
	StartTime int
	EndTime int
	//todo
}

type Position struct {
	ZoneID string
	//RC example 1-1, 2-1
	PositionID string
	// occupied, hidden, booked
	PositionStatus string
	UserID string
}
