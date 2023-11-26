package service

import (
	"crypto/sha1"
	"fmt"
	"midas"
	"midas/internal/storage"
)

const (
	salt  = "LKnxlntomLJlAsdnvizmzlsiHqp"
	salt2 = "IaqnArkj#%aInas3Inai5FSFJlasdNljaL#35Laicb3KSFjHaBiascpw"
)

type AuthService struct {
	stor storage.Authorization
}

func NewAuthService(stor storage.Authorization) *AuthService {
	return &AuthService{stor: stor}
}

func (a *AuthService) CreateUser(user *midas.User) (int64, error) {
	user.Categories = midas.StandartCategories
	user.Password = generatePasswordHash(user.Password)
	return a.stor.CreateUser(user)
}

func (a *AuthService) SignIn(login, password string) (int, error) {
	id, passw := a.stor.SignIn(login)
	if passw == generatePasswordHash(password) {
		return id, nil
	}
	return -1, fmt.Errorf("Wrong password")
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func GenerateToken(passwordHash string) string {
	hash := sha1.New()
	hash.Write([]byte(passwordHash))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt2)))
}
