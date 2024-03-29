package registry

import (
	"github.com/k-jun/waxflower/model"
)

func (sql *MySQL) SelectSeat(s *model.Seat) (*model.Seat, error) {
	err := sql.db.Get(s, "SELECT id, `row`, col, sec FROM seats WHERE id = ?", s.Id)
	return s, err
}

func (sql *MySQL) InsertSeat(s *model.Seat) (*model.Seat, error) {
	_, err := sql.db.Exec("INSERT seats(id, `row`, col, sec) VALUES(?, ?, ?, ?)",
		s.Id, s.Row, s.Col, s.Sec,
	)
	return s, err
}
