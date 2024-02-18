package util

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/k-jun/waxflower/model"
)

func InsertUser(t *testing.T, db *sqlx.DB, user *model.User) *model.User {
	_, err := db.Exec("INSERT users(id, email) VALUES(?, ?)",
		user.Id, user.Email,
	)
	if err != nil {
		t.Fatal(err)
	}
	return user
}

func SelectUser(t *testing.T, db *sqlx.DB, user *model.User) *model.User {
	err := db.Get(user, "SELECT id, email FROM users WHERE id = ?", user.Id)
	if err != nil {
		t.Fatal(err)
	}
	return user
}

func DeleteAll(db *sqlx.DB) {
	_ = db.MustExec("DELETE FROM users")
	_ = db.MustExec("DELETE FROM games")
	_ = db.MustExec("DELETE FROM seats")
	_ = db.MustExec("DELETE FROM tickets")
}
