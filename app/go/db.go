package openapi

import (
	"errors"
	"sync"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrAlreadyExist = errors.New("already exist")
)

type IDB interface {
	SelectUser(*User) (*User, error)
	InsertUser(*User) (*User, error)
	UpdateUser(*User) (*User, error)
	DeleteUser(*User) (*User, error)
	SelectGame(*Game) (*Game, error)
	InsertGame(*Game) (*Game, error)
	UpdateGame(*Game) (*Game, error)
	DeleteGame(*Game) (*Game, error)
	SelectSeat(*Seat) (*Seat, error)
	InsertSeat(*Seat) (*Seat, error)
	UpdateSeat(*Seat) (*Seat, error)
	DeleteSeat(*Seat) (*Seat, error)
	SelectTicket(*Ticket) (*Ticket, error)
	InsertTicket(*Ticket) (*Ticket, error)
	UpdateTicket(*Ticket) (*Ticket, error)
	DeleteTicket(*Ticket) (*Ticket, error)
	BuyTicket(*Ticket) (*Ticket, error)
	SearchTickets(
		dateBefore string,
		dateAfter string,
		secFrom int32,
		secTo int32,
		priceMin int32,
		priceMax int32,
	) ([]*Ticket, error)
}

type MemoryDB struct {
	users   map[string]User
	games   map[string]Game
	seats   map[string]Seat
	tickets map[string]Ticket
	rw      sync.RWMutex
}

func NewMemoryDB() IDB {
	return &MemoryDB{
		users:   map[string]User{},
		games:   map[string]Game{},
		seats:   map[string]Seat{},
		tickets: map[string]Ticket{},
	}
}

func (db *MemoryDB) SelectUser(u *User) (*User, error) {
	db.rw.RLock()
	defer db.rw.RUnlock()
	v, ok := db.users[u.Id]
	if !ok {
		return nil, ErrNotFound
	}
	return &v, nil
}

func (db *MemoryDB) InsertUser(u *User) (*User, error) {
	db.rw.Lock()
	defer db.rw.Unlock()
	_, ok := db.users[u.Id]
	if ok {
		return nil, ErrAlreadyExist
	}
	db.users[u.Id] = *u
	return u, nil
}

func (db *MemoryDB) UpdateUser(u *User) (*User, error) {
	db.rw.Lock()
	defer db.rw.Unlock()
	_, ok := db.users[u.Id]
	if !ok {
		return nil, ErrNotFound
	}
	db.users[u.Id] = *u
	return u, nil
}

func (db *MemoryDB) DeleteUser(u *User) (*User, error) {
	db.rw.Lock()
	defer db.rw.Unlock()
	_, ok := db.users[u.Id]
	if !ok {
		return nil, ErrNotFound
	}
	delete(db.users, u.Id)
	return u, nil
}

func (db *MemoryDB) SelectGame(g *Game) (*Game, error) {
	db.rw.RLock()
	defer db.rw.RUnlock()
	v, ok := db.games[g.Id]
	if !ok {
		return nil, ErrNotFound
	}
	return &v, nil
}

func (db *MemoryDB) InsertGame(g *Game) (*Game, error) {
	db.rw.Lock()
	defer db.rw.Unlock()
	_, ok := db.games[g.Id]
	if ok {
		return nil, ErrAlreadyExist
	}
	db.games[g.Id] = *g
	return g, nil
}

func (db *MemoryDB) UpdateGame(g *Game) (*Game, error) {
	db.rw.Lock()
	defer db.rw.Unlock()
	_, ok := db.games[g.Id]
	if !ok {
		return nil, ErrNotFound
	}
	db.games[g.Id] = *g
	return g, nil
}

func (db *MemoryDB) DeleteGame(g *Game) (*Game, error) {
	db.rw.Lock()
	defer db.rw.Unlock()
	_, ok := db.games[g.Id]
	if !ok {
		return nil, ErrNotFound
	}
	delete(db.games, g.Id)
	return g, nil
}

func (db *MemoryDB) SelectSeat(s *Seat) (*Seat, error) {
	db.rw.RLock()
	defer db.rw.RUnlock()
	v, ok := db.seats[s.Id]
	if !ok {
		return nil, ErrNotFound
	}
	return &v, nil
}

func (db *MemoryDB) InsertSeat(s *Seat) (*Seat, error) {
	db.rw.Lock()
	defer db.rw.Unlock()
	_, ok := db.seats[s.Id]
	if ok {
		return nil, ErrAlreadyExist
	}
	db.seats[s.Id] = *s
	return s, nil
}

func (db *MemoryDB) UpdateSeat(s *Seat) (*Seat, error) {
	db.rw.Lock()
	defer db.rw.Unlock()
	_, ok := db.seats[s.Id]
	if !ok {
		return nil, ErrNotFound
	}
	db.seats[s.Id] = *s
	return s, nil
}

func (db *MemoryDB) DeleteSeat(s *Seat) (*Seat, error) {
	db.rw.Lock()
	defer db.rw.Unlock()
	_, ok := db.seats[s.Id]
	if !ok {
		return nil, ErrNotFound
	}
	delete(db.seats, s.Id)
	return s, nil
}

func (db *MemoryDB) SelectTicket(t *Ticket) (*Ticket, error) {
	db.rw.RLock()
	defer db.rw.RUnlock()
	v, ok := db.tickets[t.Id]
	if !ok {
		return nil, ErrNotFound
	}
	return &v, nil
}

func (db *MemoryDB) InsertTicket(t *Ticket) (*Ticket, error) {
	db.rw.Lock()
	defer db.rw.Unlock()
	_, ok := db.tickets[t.Id]
	if ok {
		return nil, ErrAlreadyExist
	}
	db.tickets[t.Id] = *t
	return t, nil
}

func (db *MemoryDB) UpdateTicket(t *Ticket) (*Ticket, error) {
	db.rw.Lock()
	defer db.rw.Unlock()
	_, ok := db.tickets[t.Id]
	if !ok {
		return nil, ErrNotFound
	}
	db.tickets[t.Id] = *t
	return t, nil
}

func (db *MemoryDB) DeleteTicket(t *Ticket) (*Ticket, error) {
	db.rw.Lock()
	defer db.rw.Unlock()
	_, ok := db.tickets[t.Id]
	if !ok {
		return nil, ErrNotFound
	}
	delete(db.tickets, t.Id)
	return t, nil
}

func (db *MemoryDB) BuyTicket(t *Ticket) (*Ticket, error) {
	db.rw.Lock()
	defer db.rw.Unlock()
	v, ok := db.tickets[t.Id]
	if !ok {
		return nil, ErrNotFound
	}
	v.UserId = t.UserId
	db.tickets[t.Id] = v
	return &v, nil
}

func (db *MemoryDB) SearchTickets(
	dateBefore string,
	dateAfter string,
	secFrom int32,
	secTo int32,
	priceMin int32,
	priceMax int32,
) ([]*Ticket, error) {
	db.rw.RLock()
	defer db.rw.RUnlock()

	ts := []*Ticket{}

	for _, v := range db.tickets {
		if v.Price < int64(priceMin) || v.Price > int64(priceMax) {
			continue
		}
		sv := db.seats[v.SeatId]
		if sv.Sec < secFrom || sv.Sec > secTo {
			continue
		}
		gv := db.games[v.GameId]
		if gv.Date < dateAfter || gv.Date > dateBefore {
			continue
		}
		ts = append(ts, &v)
	}
	return ts, nil
}
