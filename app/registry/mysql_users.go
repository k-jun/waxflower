package registry

import (
	"github.com/jmoiron/sqlx"
	"github.com/k-jun/waxflower/model"
)

type MySQL struct {
	db *sqlx.DB
}

func NewMySQL(db *sqlx.DB) IRegistry {
	return &MySQL{db}
}

func (sql *MySQL) SelectUser(u *model.User) (*model.User, error) {
	err := sql.db.Get(u, "SELECT id, email FROM users WHERE id = ?", u.Id)
	return u, err
}

func (sql *MySQL) InsertUser(_ *model.User) (*model.User, error) {
	panic("not implemented") // TODO: Implement
}

func (sql *MySQL) UpdateUser(_ *model.User) (*model.User, error) {
	panic("not implemented") // TODO: Implement
}

func (sql *MySQL) DeleteUser(_ *model.User) (*model.User, error) {
	panic("not implemented") // TODO: Implement
}
