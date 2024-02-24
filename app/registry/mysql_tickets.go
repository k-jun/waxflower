package registry

import (
	"github.com/k-jun/waxflower/model"
)

func (sql *MySQL) SelectTicket(t *model.Ticket) (*model.Ticket, error) {
	err := sql.db.Get(t, `
SELECT 
  tickets.id as "id", tickets.price as "price",
	seats.id as "seat.id", seats.sec as "seat.sec", seats.col as "seat.col", seats.row as "seat.row",
	games.id as "game.id", games.date as "game.date", games.time as "game.time"
FROM tickets
	LEFT JOIN seats ON tickets.seat_id = seats.id
	LEFT JOIN games ON tickets.game_id = games.id
WHERE tickets.id = ?`, t.Id)
	return t, err
}

func (sql *MySQL) InsertTicket(t *model.Ticket) (*model.Ticket, error) {
	_, err := sql.db.Exec("INSERT tickets(id, price, game_id, seat_id) VALUES(?, ?, ?, ?)",
		t.Id, t.Price, t.Game.Id, t.Seat.Id,
	)
	return t, err
}
