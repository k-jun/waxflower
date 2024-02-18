package registry

import (
	"github.com/jmoiron/sqlx"
	"github.com/k-jun/waxflower/model"
)

type IRegistry interface {
	SelectUser(*model.User) (*model.User, error)
	InsertUser(*model.User) (*model.User, error)
	UpdateUser(*model.User) (*model.User, error)
	DeleteUser(*model.User) (*model.User, error)

	SelectGame(*model.Game) (*model.Game, error)
	InsertGame(*model.Game) (*model.Game, error)
	UpdateGame(*model.Game) (*model.Game, error)
	DeleteGame(*model.Game) (*model.Game, error)

	SelectSeat(*model.Seat) (*model.Seat, error)
	InsertSeat(*model.Seat) (*model.Seat, error)
	UpdateSeat(*model.Seat) (*model.Seat, error)
	DeleteSeat(*model.Seat) (*model.Seat, error)

	SelectTicket(*model.Ticket) (*model.Ticket, error)
	InsertTicket(*model.Ticket) (*model.Ticket, error)
	UpdateTicket(*model.Ticket) (*model.Ticket, error)
	DeleteTicket(*model.Ticket) (*model.Ticket, error)

	Reset() error
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
