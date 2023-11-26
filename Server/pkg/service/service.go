package service

import (
	"midas"
	"midas/internal/storage"
)

type Authorization interface {
	CreateUser(user *midas.User) (int64, error)
	SignIn(login, password string) (int, error)
}

type User interface {
	GetCategories(user *midas.UserData) []midas.Category
	AddCategory(categoryRequest *midas.CategoryRequest) (int, error)
	MakeTransaction(categoryRequest *midas.CategoryRequest) error
	GetUser(id int64) (midas.User, error)
	CheckUser(id int, token string) bool
	GetMainData(id int64, timeStart, timeEnd int) *midas.MainData
	DeleteCategory(user *midas.UserData, category *midas.Category) error
}

type Service struct {
	Authorization
	User
}

func NewService(db *storage.Storage) *Service {
	return &Service{
		Authorization: NewAuthService(db.Authorization),
		User:          NewUserService(db.User),
	}
}
