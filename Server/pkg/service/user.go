package service

import (
	"log"
	"midas"
	"midas/internal/storage"
)

type UserService struct {
	stor storage.User
}

func NewUserService(stor storage.User) *UserService {
	return &UserService{stor: stor}
}

func (c *UserService) NewCategory(id int, ctgr midas.Category) error {
	return nil
}

func (c *UserService) GetCategories(id int) []midas.Category {
	log.Print("Service: GetCategory")
	return c.stor.GetCategories(id)
}
