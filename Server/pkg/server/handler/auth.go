package handler

import (
	"fmt"
	"log"
	"midas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")

	var newUser midas.User

	if err := c.BindJSON(&newUser); err != nil {
		log.Fatal(err)
	}

	id, err := h.service.Authorization.CreateUser(newUser)
	if err != nil {
		log.Fatalf("Handler Error: %v", err)
	}

	log.Printf("New User: {id: %v, name: %v, login: %v}", id, newUser.Name, newUser.Login)

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("userId", fmt.Sprint(id), 3600*24*30, "/", "", false, true)

	c.Status(http.StatusOK)
}

func (h *Handler) signIn(c *gin.Context) {
	c.Status(http.StatusBadRequest)
}
