package storage

import (
	"database/sql"
	"log"
	"midas"
	"strconv"
	"strings"
)

type UserDB struct {
	db *sql.DB
}

func NewUserDB(db *sql.DB) *UserDB {
	return &UserDB{db: db}
}

func (c *UserDB) NewCategory(id int, catg midas.Category) error {
	return nil
}

func (c *UserDB) GetCategories(id int) []midas.Category {
	log.Printf("Sqlite: GetCategories")
	rows, _ := c.db.Query("SELECT categories FROM users WHERE id=?", id)
	defer rows.Close()

	var catg string
	for rows.Next() {
		rows.Scan(&catg)
	}
	catg = catg[1 : len([]rune(catg))-1]
	sliceCatg := strings.Split(strings.Join(strings.Split(catg, "{"), ""), "} ")
	answ := []midas.Category{}
	for _, val := range sliceCatg {
		nVal := strings.Split(val, " ")
		idCatg, _ := strconv.Atoi(nVal[0])
		amountLimit, _ := strconv.Atoi(nVal[2])
		amount, _ := strconv.Atoi(nVal[3])
		answ = append(answ, midas.Category{Id: idCatg, Name: nVal[1], AmountLimit: amountLimit, Amount: amount})
	}

	return answ
}
