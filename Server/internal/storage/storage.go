package storage

import (
	"database/sql"
	"midas"
)

type Authorization interface {
	CreateUser(*midas.User) (int64, error)
	SignIn(login string) (int, string)
}

type User interface {
	GetCategories(id int64) []midas.Category
	AddCategory(categoryRequest *midas.CategoryRequest) (int, error)
	GetUser(id int64) (midas.User, error)
	MakeTransaction(userId int64, categoryId, amount int) error
	CalculationOfExpenses(timeStart, timeEnd int) int
	DeleteCategory(userId int64, categoryId, amount int) error
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
