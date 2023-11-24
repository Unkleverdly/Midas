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

func (a *Auth) CreateUser(user midas.User) (int64, error) {
	log.Print("Database Auth")
	log.Print(user)
	db, _ := a.db.Prepare("INSERT INTO users(name, login, password, categories) VALUES(?, ?, ?, ?)")
	defer db.Close()
	res, err := db.Exec(user.Name, user.Login, user.Password, fmt.Sprint(user.Categories))
	id, _ := res.LastInsertId()
	if err != nil {
		log.Fatalf("Database error: %v", err)
	}
	return id, nil
}
