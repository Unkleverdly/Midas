package service

import (
	"midas"
	"midas/internal/storage"
)

type Authorization interface {
	CreateUser(midas.User) (int64, error)
}

type User interface {
	NewCategory(int, midas.Category) error
	GetCategories(int) []midas.Category
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
