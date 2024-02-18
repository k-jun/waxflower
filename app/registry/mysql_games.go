package registry

import (
	"github.com/k-jun/waxflower/model"
)

func (sql *MySQL) SelectGame(g *model.Game) (*model.Game, error) {
	err := sql.db.Get(g, "SELECT id, date, time FROM games WHERE id = ?", g.Id)
	return g, err
}

func (sql *MySQL) InsertGame(g *model.Game) (*model.Game, error) {
	_, err := sql.db.Exec("INSERT Games(id, email) VALUES(?, ?)") // u.Id, u.Email,

	return g, err
}

func (sql *MySQL) UpdateGame(_ *model.Game) (*model.Game, error) {
	panic("not implemented") // TODO: Implement
}

func (sql *MySQL) DeleteGame(_ *model.Game) (*model.Game, error) {
	panic("not implemented") // TODO: Implement
}
