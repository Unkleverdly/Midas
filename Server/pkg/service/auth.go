package service

import (
	"log"
	"midas"
	"midas/internal/storage"
)

type AuthService struct {
	stor storage.Authorization
}

func NewAuthService(stor storage.Authorization) *AuthService {
	return &AuthService{stor: stor}
}

func (a *AuthService) CreateUser(user midas.User) (int64, error) {
	user.Categories = midas.StandartCategories
	log.Print("Service Auth")
	return a.stor.CreateUser(user)
}
