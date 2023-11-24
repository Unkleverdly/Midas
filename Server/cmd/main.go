package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"midas/internal/storage"
	"midas/pkg/server"
	"midas/pkg/server/handler"
	"midas/pkg/service"
)

func main() {
	database, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}

	db, _ := storage.NewStorage(database)
	services := service.NewService(db)
	handl := handler.NewHandler(services)
	serv := new(server.Server)

	if err := serv.Run("8080", handl.Init()); err != nil {
		log.Fatal(err)
	}

}
