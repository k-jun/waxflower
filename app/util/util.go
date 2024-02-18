package util

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/k-jun/waxflower/model"
)

func InsertUser(t *testing.T, db *sqlx.DB, u *model.User) *model.User {
	_, err := db.Exec("INSERT users(id, email) VALUES(?, ?)",
		u.Id, u.Email,
	)
	if err != nil {
		t.Fatal(err)
	}
	return u
}

func SelectUser(t *testing.T, db *sqlx.DB, u *model.User) *model.User {
	err := db.Get(u, "SELECT id, email FROM users WHERE id = ?", u.Id)
	if err != nil {
		t.Fatal(err)
	}
	return u
}

func InsertGame(t *testing.T, db *sqlx.DB, g *model.Game) *model.Game {
	_, err := db.Exec("INSERT games(id, date, time) VALUES(?, ?, ?)",
		g.Id, g.Date, g.Time,
	)
	if err != nil {
		t.Fatal(err)
	}
	return g
}

func SelectGame(t *testing.T, db *sqlx.DB, g *model.Game) *model.Game {
	err := db.Get(g, "SELECT id, date, time FROM games WHERE id = ?", g.Id)
	if err != nil {
		t.Fatal(err)
	}
	return g
}

func InsertSeat(t *testing.T, db *sqlx.DB, s *model.Seat) *model.Seat {
	_, err := db.Exec("INSERT seats(id, `row`, col, sec) VALUES(?, ?, ?, ?)",
		s.Id, s.Row, s.Col, s.Sec,
	)
	if err != nil {
		t.Fatal(err)
	}
	return s
}

func SelectSeat(t *testing.T, db *sqlx.DB, s *model.Seat) *model.Seat {
	err := db.Get(s, "SELECT id, `row`, col, sec FROM seats WHERE id = ?", s.Id)
	if err != nil {
		t.Fatal(err)
	}
	return s
}

func InsertTicket(t *testing.T, db *sqlx.DB, ti *model.Ticket) *model.Ticket {
	_, err := db.Exec("INSERT tickets(id, price, user_id, game_id, seat_id) VALUES(?, ?, ?, ?, ?)",
		ti.Id, ti.Price, ti.UserId, ti.GameId, ti.SeatId,
	)
	if err != nil {
		t.Fatal(err)
	}
	return ti
}

func SelectTicket(t *testing.T, db *sqlx.DB, ti *model.Ticket) *model.Ticket {
	_, err := db.Exec("INSERT tickets(id, price, user_id, game_id, seat_id) VALUES(?, ?, ?, ?, ?)",
		ti.Id, ti.Price, ti.UserId, ti.GameId, ti.SeatId,
	)
	if err != nil {
		t.Fatal(err)
	}
	return ti
}

func DeleteAll(db *sqlx.DB) {
	_ = db.MustExec("DELETE FROM tickets")
	_ = db.MustExec("DELETE FROM users")
	_ = db.MustExec("DELETE FROM games")
	_ = db.MustExec("DELETE FROM seats")
}
