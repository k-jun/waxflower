package registry

import (
	"github.com/jmoiron/sqlx"
	"github.com/k-jun/waxflower/model"
)

type IRegistry interface {
	SelectUser(*model.User) (*model.User, error)
	InsertUser(*model.User) (*model.User, error)

	SelectGame(*model.Game) (*model.Game, error)
	InsertGame(*model.Game) (*model.Game, error)

	SelectSeat(*model.Seat) (*model.Seat, error)
	InsertSeat(*model.Seat) (*model.Seat, error)

	SelectTicket(*model.Ticket) (*model.Ticket, error)
	InsertTicket(*model.Ticket) (*model.Ticket, error)

	Reset() error
	Search(*model.Ticket) ([]*model.Ticket, error)
	Reserve(*model.Ticket) (*model.Ticket, error)
}

type MySQL struct {
	db *sqlx.DB
}

func NewMySQL(db *sqlx.DB) IRegistry {
	return &MySQL{db}
}

func (sql *MySQL) Reset() error {
	tx, err := sql.db.Beginx()
	defer tx.Commit()
	if err != nil {
		return err
	}
	_, err = sql.db.Exec("DELETE FROM tickets")
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	_, err = tx.Exec("DELETE FROM users")
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	_, err = tx.Exec("DELETE FROM games")
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	_, err = tx.Exec("DELETE FROM seats")
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return err
}

func (sql *MySQL) Search(t *model.Ticket) ([]*model.Ticket, error) {
	args := []any{t.Game.Date, t.Seat.Sec}
	rows, err := sql.db.Queryx(`
SELECT
  tickets.id as "id", tickets.price as "price",
	seats.id as "seat.id", seats.sec as "seat.sec", seats.col as "seat.col", seats.row as "seat.row",
	games.id as "game.id", games.date as "game.date", games.time as "game.time"
FROM tickets
	LEFT JOIN seats ON tickets.seat_id = seats.id
	LEFT JOIN games ON tickets.game_id = games.id
	LEFT JOIN users ON tickets.user_id = users.id
WHERE games.date = ? AND seats.sec = ?`, args...)
	ts := []*model.Ticket{}
	if err != nil {
		return ts, err
	}

	for rows.Next() {
		nt := &model.Ticket{}
		if err := rows.StructScan(nt); err != nil {
			return ts, err
		}
		ts = append(ts, nt)
	}
	return ts, err
}

func (sql *MySQL) Reserve(_ *model.Ticket) (*model.Ticket, error) {
	panic("not implemented") // TODO: Implement
}
