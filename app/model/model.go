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
	UserId string `db:"user_id"`
	GameId string `db:"game_id"`
	SeatId string `db:"seat_id"`
}
type User struct {
	Id    string
	Email string
}
