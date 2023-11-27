package storage

import (
	"database/sql"
	"fmt"
	"log"
	"midas"
	"strconv"
	"strings"
	"time"
)

type UserDB struct {
	db *sql.DB
}

func NewUserDB(db *sql.DB) *UserDB {
	return &UserDB{db: db}
}

func (u *UserDB) GetCategories(id int64, time ...int) []midas.Category {
	rows, _ := u.db.Query("SELECT categories FROM users WHERE id=?", id)
	defer rows.Close()

	var catg string
	for rows.Next() {
		rows.Scan(&catg)
	}

	catg = catg[1 : len([]rune(catg))-1]

	sliceCatg := strings.Split(strings.Join(strings.Split(catg, "{"), ""), "} ")
	answ := []midas.Category{}
	for i, val := range sliceCatg {
		var amount int
		nVal := strings.Split(val, " ")
		idCatg, _ := strconv.Atoi(nVal[0])
		amountLimit, _ := strconv.Atoi(nVal[2])
		if len(time) != 0 {
			rows, err := u.db.Query("SELECT amount FROM transactions WHERE time < ? AND time > ? AND categoryId = ? AND userId = ?", time[1], time[0], idCatg, id)
			if err != nil {
				log.Print("CalculationOfExpenses error: ", err)
			}

			var nums []int
			for rows.Next() {
				var num string
				rows.Scan(&num)
				numForNums, _ := strconv.Atoi(num)
				nums = append(nums, numForNums)
			}
			amount = sum(nums)
		} else if i == len(sliceCatg)-1 {
			amount, _ = strconv.Atoi(nVal[3][:len(nVal[3])-1])
		} else {
			amount, _ = strconv.Atoi(nVal[3][:len(nVal[3])])
		}
		answ = append(answ, midas.Category{Id: idCatg, Name: nVal[1], AmountLimit: amountLimit, Amount: amount})
	}

	return answ
}

func (u *UserDB) AddCategory(req *midas.CategoryRequest) (int, error) {
	var isNew bool = true
	answ := u.GetCategories(req.UserData.Id)
	var catg = req.Category
	catg.Id = answ[len(answ)-1].Id + 1

	for i, val := range answ {
		if val.Id == req.Category.Id {
			answ[i].AmountLimit = req.Category.AmountLimit
			answ[i].Name = req.Category.Name
			isNew = false
			break
		}
	}

	if isNew {
		answ = append(answ, *catg)
	}

	_, err := u.db.Exec("UPDATE users set categories = ? WHERE id = ?", fmt.Sprint(answ), req.UserData.Id)

	if err != nil {
		log.Print(err)
		return 0, err
	}
	if !isNew {
		return req.Category.Id, nil
	}
	return catg.Id, nil
}

func (u *UserDB) DeleteCategory(userId int64, categoryId, amount int) error {
	categ := u.GetCategories(userId)

	for i, val := range categ {
		if val.Id == 0 {
			categ[i].Amount += amount
		}
		if val.Id == categoryId {
			categ = append(categ[:i], categ[i+1:]...)
			break
		}
	}
	_, err := u.db.Exec("UPDATE users set categories = ? WHERE id = ?", fmt.Sprint(categ), userId)
	_, err = u.db.Exec("UPDATE transactions set categoryId = 0 WHERE userId = ? AND categoryId = ?", userId, categoryId)
	return err
}

func (u *UserDB) GetUser(id int64) (midas.User, error) {
	rows := u.db.QueryRow("SELECT name, login, hash_password FROM users WHERE id = ?", id)
	var prod struct {
		name     string
		login    string
		password string
	}
	rows.Scan(&prod.name, &prod.login, &prod.password)

	return midas.User{
		Id:         id,
		Login:      prod.login,
		Name:       prod.name,
		Password:   prod.password,
		Categories: u.GetCategories(id),
	}, nil
}

func (u *UserDB) MakeTransaction(userId int64, categoryId, amount int) error {
	var catg []midas.Category = u.GetCategories(userId)
	for i, val := range catg {
		if val.Id == categoryId {
			catg[i].Amount += amount
			break
		}
	}
	_, err := u.db.Exec("UPDATE users set categories = ? WHERE id = ?", fmt.Sprint(catg), userId)
	if err != nil {
		return err
	}
	_, err = u.db.Exec("INSERT INTO transactions(userId, categoryId, amount, time) VALUES(?, ?, ?, ?) ", userId, categoryId, amount, time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

func (u *UserDB) GetTransactions(userId int64, timeStart, timeEnd int) []midas.Transaction {
	rows, err := u.db.Query("SELECT id, categoryId, amount, time FROM transactions WHERE time < ? AND time > ? AND userId = ?", timeEnd, timeStart, userId)
	if err != nil {
		log.Print("GetTrascation error: ", err)
	}
	defer rows.Close()

	answ := []midas.Transaction{}

	for rows.Next() {
		t := midas.Transaction{}
		err := rows.Scan(&t.Id, &t.CategoryId, &t.Amount, &t.Time)
		if err != nil {
			log.Print("Database error", err)
			continue
		}

		answ = append(answ, t)
	}
	return answ
}

func (u *UserDB) CalculationOfExpenses(timeStart, timeEnd int, id int64) int {
	rows, err := u.db.Query("SELECT amount FROM transactions WHERE time < ? AND time > ? AND userId = ?", timeEnd, timeStart, id)
	if err != nil {
		log.Print("CalculationOfExpenses error: ", err)
		return 0
	}

	var nums []int
	for rows.Next() {
		var num string
		rows.Scan(&num)
		numForNums, _ := strconv.Atoi(num)
		nums = append(nums, numForNums)
	}

	return sum(nums)
}

func sum(arr []int) int {
	sum := 0
	for _, valueInt := range arr {
		sum += valueInt
	}
	return sum
}
