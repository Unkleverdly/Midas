package handler

import (
	"log"
	"midas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var newUser midas.User

	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, midas.ApiAnswer{Status: 3})
			log.Print("Error: ", err)

		}
	}()

	if err := c.BindJSON(&newUser); err != nil {
		log.Printf("JSON errror: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := h.service.Authorization.CreateUser(&newUser)

	if err != nil {
		log.Printf("Handler Error: %v", err)
		c.AbortWithStatusJSON(http.StatusOK, midas.ApiAnswer{Status: 3})
		return
	}
	newUser, _ = h.service.GetUser(id)

	log.Printf("New User: {id: %v, name: %v, login: %v}", id, newUser.Name, newUser.Login)
	c.JSON(http.StatusOK, midas.ApiAnswer{
		Status: 0,
		Result: &midas.UserData{
			Id:    newUser.Id,
			Token: newUser.Token,
		},
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var User struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&User); err != nil {
		log.Print("JSON errror: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := h.service.SignIn(User.Login, User.Password)
	if err != nil {
		log.Print("Wrong Password")
		c.AbortWithStatusJSON(http.StatusOK, midas.ApiAnswer{Status: 2})
		return
	}
	newUser, _ := h.service.GetUser(int64(id))

	c.JSON(http.StatusOK, midas.ApiAnswer{
		Status: 0,
		Result: &midas.UserData{
			Id:    newUser.Id,
			Token: newUser.Token,
		},
	})
}
