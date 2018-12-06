package zone

type zone struct {
	zoneID string
	positionNumber int
	CreateTime int
	StartTime int
	EndTime int
	//todo
}

type position struct {
	zoneID string
	//RC example A1, AA1
	positionID string
	// occupied, hidden, booked
	positionStatus string
	UserID string
}
