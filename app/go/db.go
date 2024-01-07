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
}

type MemoryDB struct {
	users map[string]User
	games map[string]Game
	seats map[string]Seat
	rw    sync.RWMutex
}

func NewMemoryDB() IDB {
	return &MemoryDB{
		users: map[string]User{},
		games: map[string]Game{},
		seats: map[string]Seat{},
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
	if _, ok := db.users[u.Id]; !ok {
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
	if _, ok := db.games[g.Id]; !ok {
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
	if _, ok := db.users[s.Id]; !ok {
		return nil, ErrNotFound
	}
	delete(db.seats, s.Id)
	return s, nil
}
