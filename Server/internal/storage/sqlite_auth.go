package storage

import (
	"database/sql"
	"fmt"
	"log"
	"midas"
)

type Auth struct {
	db *sql.DB
}

func NewAuthDB(db *sql.DB) *Auth {
	return &Auth{db: db}
}

func (a *Auth) CreateUser(user *midas.User) (int64, error) {
	res, err := a.db.Exec("INSERT INTO users(name, login, hash_password, categories) VALUES(?, ?, ?, ?)", user.Name, user.Login, user.Password, fmt.Sprint(user.Categories))

	id, _ := res.LastInsertId()
	if err != nil {
		log.Fatalf("Database error: %v", err)
	}
	return id, nil
}

func (a *Auth) SignIn(login string) (int, string) {
	rows := a.db.QueryRow(`SELECT id, hash_password FROM users WHERE login = ?`, login)

	var prod struct {
		id       int
		password string
	}

	rows.Scan(&prod.id, &prod.password)

	return prod.id, prod.password
}
