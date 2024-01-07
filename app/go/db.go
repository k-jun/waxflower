package openapi

import (
	"errors"
	"sync"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrAlreadyExist = errors.New("already exist")
)

type IUserDB interface {
	SelectUser(*User) (*User, error)
	InsertUser(*User) (*User, error)
	UpdateUser(*User) (*User, error)
	DeleteUser(*User) (*User, error)
}

type UserDB struct {
	storage map[string]User
	rw      sync.RWMutex
}

func NewUserDB() IUserDB {
	return &UserDB{}
}

func (db *UserDB) SelectUser(u *User) (*User, error) {
	db.rw.RLock()
	defer db.rw.RUnlock()
	if v, ok := db.storage[u.Id]; ok {
		return &v, nil
	}
	return nil, ErrNotFound
}

func (db *UserDB) InsertUser(u *User) (*User, error) {
	db.rw.Lock()
	defer db.rw.Unlock()
	if _, ok := db.storage[u.Id]; !ok {
		return nil, ErrAlreadyExist
	}
	db.storage[u.Id] = *u
	return u, nil
}

func (db *UserDB) UpdateUser(u *User) (*User, error) {
	db.rw.Lock()
	defer db.rw.Unlock()
	if _, ok := db.storage[u.Id]; ok {
		db.storage[u.Id] = *u
		return u, nil
	}
	return nil, ErrNotFound
}

func (db *UserDB) DeleteUser(u *User) (*User, error) {
	db.rw.Lock()
	defer db.rw.Unlock()
	if _, ok := db.storage[u.Id]; ok {
		delete(db.storage, u.Id)
	}
	return nil, ErrNotFound
}