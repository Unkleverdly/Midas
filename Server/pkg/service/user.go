package service

import (
	"midas"
	"midas/internal/storage"
)

type UserService struct {
	stor storage.User
}

func NewUserService(stor storage.User) *UserService {
	return &UserService{stor: stor}
}

func (u *UserService) GetCategories(req *midas.UserData) []midas.Category {
	return u.stor.GetCategories(req.Id)
}

func (u *UserService) AddCategory(req *midas.CategoryRequest) (int, error) {
	return u.stor.AddCategory(req)
}

func (u *UserService) GetUser(id int64) (midas.User, error) {
	user, err := u.stor.GetUser(id)
	user.Token = GenerateToken(user.Password)
	return user, err
}

func (u *UserService) GetMainData(id int64, timeStart, timeEnd int) *midas.MainData {
	user, _ := u.GetUser(id)
	return &midas.MainData{
		Name:           user.Name,
		Categories:     user.Categories,
		MonthSpendings: u.stor.CalculationOfExpenses(timeStart, timeEnd),
	}
}

func (u *UserService) DeleteCategory(user *midas.UserData, category *midas.Category) error {
	return u.stor.DeleteCategory(user.Id, category.Id, category.Amount)
}

func (u *UserService) MakeTransaction(req *midas.CategoryRequest) error {
	return u.stor.MakeTransaction(req.UserData.Id, req.Category.Id, req.Category.Amount)
}

func (u *UserService) CheckUser(id int, token string) bool {
	user, _ := u.stor.GetUser(int64(id))
	if GenerateToken(user.Password) == token {
		return true
	}
	return false
}
