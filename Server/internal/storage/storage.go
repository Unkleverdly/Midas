package storage

import (
	"database/sql"
	"midas"
)

type Authorization interface {
	CreateUser(midas.User) (int64, error)
}

type User interface {
	NewCategory(int, midas.Category) error
	GetCategories(int) []midas.Category
}

type Storage struct {
	Authorization
	User
}

func NewStorage(db *sql.DB) (*Storage, error) {
	return &Storage{
		Authorization: NewAuthDB(db),
		User:          NewUserDB(db),
	}, nil
}
