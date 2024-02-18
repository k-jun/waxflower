package model

type Game struct {
	Id   string
	Date string
	Time string
}
type Seat struct {
	Id  string
	Row string
	Col int32
	Sec int32
}
type Ticket struct {
	Id     string
	Price  int64
	UserId string
	GameId string
	SeatId string
}
type User struct {
	Id    string
	Email string
}
