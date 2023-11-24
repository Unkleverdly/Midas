package midas

type User struct {
	Id         int64  `json:"id"`
	Name       string `json:"name" binding:"required"`
	Login      string `json:"login" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Token      string
	Categories []Category `json:"categories"`
}

type Category struct {
	Id          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	AmountLimit int    `json:"amountLimit" binding:"required"`
	Amount      int
}

var StandartCategories = []Category{
	{Id: 0, Name: "Undefined"},
	{Id: 1, Name: "Supermarket"},
	{Id: 2, Name: "Services"},
	{Id: 3, Name: "Utilities"}}
