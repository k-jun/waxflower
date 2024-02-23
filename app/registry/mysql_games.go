package registry

import (
	"github.com/k-jun/waxflower/model"
)

func (sql *MySQL) SelectGame(g *model.Game) (*model.Game, error) {
	err := sql.db.Get(g, "SELECT id, date, time FROM games WHERE id = ?", g.Id)
	return g, err
}

func (sql *MySQL) InsertGame(g *model.Game) (*model.Game, error) {
	_, err := sql.db.Exec("INSERT games(id, date, time) VALUES(?, ?, ?)",
		g.Id, g.Date, g.Time,
	)
	return g, err
}
