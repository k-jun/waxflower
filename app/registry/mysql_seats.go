package registry

import (
	"github.com/k-jun/waxflower/model"
)

func (sql *MySQL) SelectSeat(s *model.Seat) (*model.Seat, error) {
	err := sql.db.Get(s, "SELECT id, col, `row`, sec FROM seats WHERE id = ?", s.Id)
	return s, err
}

func (sql *MySQL) InsertSeat(s *model.Seat) (*model.Seat, error) {
	_, err := sql.db.Exec("INSERT seats(id, date, time) VALUES(?, ?, ?)")
	return s, err
}

func (sql *MySQL) UpdateSeat(_ *model.Seat) (*model.Seat, error) {
	panic("not implemented") // TODO: Implement
}

func (sql *MySQL) DeleteSeat(_ *model.Seat) (*model.Seat, error) {
	panic("not implemented") // TODO: Implement
}
