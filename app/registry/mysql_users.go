package registry

import (
	"github.com/k-jun/waxflower/model"
)

func (sql *MySQL) SelectUser(u *model.User) (*model.User, error) {
	err := sql.db.Get(u, "SELECT id, email FROM users WHERE id = ?", u.Id)
	return u, err
}

func (sql *MySQL) InsertUser(u *model.User) (*model.User, error) {
	_, err := sql.db.Exec("INSERT users(id, email) VALUES(?, ?)",
		u.Id, u.Email,
	)
	return u, err
}

func (sql *MySQL) UpdateUser(_ *model.User) (*model.User, error) {
	panic("not implemented") // TODO: Implement
}

func (sql *MySQL) DeleteUser(_ *model.User) (*model.User, error) {
	panic("not implemented") // TODO: Implement
}
