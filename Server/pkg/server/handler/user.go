package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Request struct {
	UserId      int
	Name        string
	DayLimit    int
	AmountLimit int
}

func (h *Handler) NewCategory(c *gin.Context) {
	var req Request
	if err := c.BindJSON(&req); err != nil {
		log.Print(err)
	}
	log.Print(req)
}

func (h *Handler) getCategories(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	defer func() {
		if err := recover(); err != nil {
			log.Print(err)
			c.Status(http.StatusBadRequest)
		}
	}()
	type shortCategory struct {
		Name   string `json:"name"`
		Amount int    `json:"amount"`
		Limit  int    `json:"limit"`
	}

	// var req struct {
	// 	Id int `json:"id"`
	// }
	idCookie, err := c.Cookie("userId")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Cookie userId: %v", idCookie)
	id, _ := strconv.Atoi(idCookie)
	// if err := c.BindJSON(&req); err != nil {
	// 	log.Print(err)
	// 	c.Status(http.StatusBadRequest)
	// 	return
	// }
	answ := h.service.User.GetCategories(id)
	var response map[int]shortCategory = make(map[int]shortCategory)

	for _, val := range answ {
		response[val.Id] = shortCategory{
			Name:   val.Name,
			Amount: val.Amount,
			Limit:  val.AmountLimit,
		}
	}

	c.JSON(http.StatusOK, response)
}
