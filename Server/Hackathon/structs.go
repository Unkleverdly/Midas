package midas

type ApiAnswer struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

type User struct {
	Id         int64      `json:"id"`
	Name       string     `json:"name"`
	Login      string     `json:"login"`
	Password   string     `json:"password"`
	Token      string     `json:"token"`
	Categories []Category `json:"categories"`
}

type Transaction struct {
	Id         int `json:"id"`
	CategoryId int `json:"categoryId"`
	Amount     int `json:"amount"`
	Time       int `json:"time"`
}

type MainData struct {
	Name           string        `json:"name"`
	MonthSpendings int           `json:"monthSpendings"`
	Categories     []Category    `json:"categories"`
	Transactions   []Transaction `json:"transactions"`
}

type UserData struct {
	Id    int64  `json:"id" binding:"required"`
	Token string `json:"token" binding:"required"`
}

type Category struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	AmountLimit int    `json:"limit"`
	Amount      int    `json:"amount"`
}

type CategoryRequest struct {
	Category *Category `json:"request"`
	UserData *UserData `json:"user"`
}

var StandartCategories = []Category{
	{Id: 0, Name: "Undefined"},
	{Id: 1, Name: "Supermarket"},
	{Id: 2, Name: "Services"},
	{Id: 3, Name: "Utilities"}}

var Status map[int]string = map[int]string{
	0: "OK",
	1: "User not found",
	2: "Wrong Paswword",
	3: "User already exist",
}
