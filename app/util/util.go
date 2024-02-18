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

func DeleteAll(db *sqlx.DB) {
	_ = db.MustExec("DELETE FROM users")
	_ = db.MustExec("DELETE FROM games")
	_ = db.MustExec("DELETE FROM seats")
	_ = db.MustExec("DELETE FROM tickets")
}
