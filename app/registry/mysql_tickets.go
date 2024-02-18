package registry

import (
	"github.com/k-jun/waxflower/model"
)

func (sql *MySQL) SelectTicket(t *model.Ticket) (*model.Ticket, error) {
	err := sql.db.Get(t, "SELECT id, price, user_id, game_id, seat_id FROM tickets WHERE id = ?", t.Id)
	return t, err
}

func (sql *MySQL) InsertTicket(t *model.Ticket) (*model.Ticket, error) {
	_, err := sql.db.Exec("INSERT tickets(id, price, user_id, game_id, seat_id) VALUES(?, ?, ?, ?, ?)",
		t.Id, t.Price, t.UserId, t.GameId, t.SeatId,
	)
	return t, err
}

func (sql *MySQL) UpdateTicket(_ *model.Ticket) (*model.Ticket, error) {
	panic("not implemented") // TODO: Implement
}

func (sql *MySQL) DeleteTicket(_ *model.Ticket) (*model.Ticket, error) {
	panic("not implemented") // TODO: Implement
}
