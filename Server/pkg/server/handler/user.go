package handler

import (
	"log"
	"midas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addCategory(c *gin.Context) {
	var req midas.CategoryRequest

	if err := c.BindJSON(&req); err != nil {
		log.Printf("JSON error: %v", err)
		c.Status(http.StatusBadRequest)
		return
	}

	if h.service.CheckUser(int(req.UserData.Id), req.UserData.Token) == false {
		log.Print("Wrong Token")
		c.JSON(http.StatusOK, midas.ApiAnswer{Status: 1})
		return
	}

	id, err := h.service.User.AddCategory(&req)
	if err != nil {
		log.Print("AddCategory error: ", err)
		return
	}

	c.JSON(http.StatusOK, map[string]int{"status": 0, "result": id})
}

func (h *Handler) getCategories(c *gin.Context) {
	var newReq midas.UserData

	if err := c.BindJSON(&newReq); err != nil {
		log.Printf("JSON error: %v", err)
		c.Status(http.StatusBadRequest)
		return
	}

	if h.service.CheckUser(int(newReq.Id), newReq.Token) == false {
		log.Print("Wrong Token")
		c.JSON(http.StatusOK, midas.ApiAnswer{Status: 1})
		return
	}

	answ := h.service.User.GetCategories(&newReq)
	c.JSON(http.StatusOK, midas.ApiAnswer{Status: 0, Result: &answ})
}

func (h *Handler) makeTransaction(c *gin.Context) {
	var newReq midas.CategoryRequest

	if err := c.BindJSON(&newReq); err != nil {
		log.Print("JSON error: ", err)
		c.Status(http.StatusBadRequest)
		return
	}
	log.Printf("Make transaction for %v id", newReq.UserData.Id)

	if h.service.CheckUser(int(newReq.UserData.Id), newReq.UserData.Token) == false {
		log.Print("Wrong Token")
		c.JSON(http.StatusOK, midas.ApiAnswer{Status: 1})
		return
	}

	err := h.service.MakeTransaction(&newReq)
	if err != nil {
		log.Print("Transaction error: ", err)
		return
	}
	c.JSON(http.StatusOK, midas.ApiAnswer{Status: 0})
}

func (h *Handler) getMainData(c *gin.Context) {
	var newReq struct {
		midas.UserData `json:"user"`
		Time           struct {
			TimeStart int `json:"timeStart"`
			TimeEnd   int `json:"timeEnd"`
		} `json:"request"`
	}

	if err := c.BindJSON(&newReq); err != nil {
		log.Printf("JSON error: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	log.Printf("Get Main Data for %v id", newReq.Id)

	if h.service.CheckUser(int(newReq.UserData.Id), newReq.UserData.Token) == false {
		log.Print("Wrong Token")
		c.AbortWithStatusJSON(http.StatusOK, midas.ApiAnswer{Status: 1})
		return
	}
	answ := h.service.GetMainData(newReq.UserData.Id, newReq.Time.TimeStart, newReq.Time.TimeEnd)

	c.JSON(http.StatusOK, midas.ApiAnswer{Status: 0, Result: *answ})
}

func (h *Handler) deleteCategory(c *gin.Context) {
	var newReq midas.CategoryRequest

	if err := c.BindJSON(&newReq); err != nil {
		log.Printf("JSON error: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	log.Printf("Delete category(userId:%v, categoryId:%v)", newReq.UserData.Id, newReq.Category.Id)

	if h.service.CheckUser(int(newReq.UserData.Id), newReq.UserData.Token) == false {
		log.Print("Wrong Token")
		c.AbortWithStatusJSON(http.StatusOK, midas.ApiAnswer{Status: 1})
		return
	}

	err := h.service.DeleteCategory(newReq.UserData, newReq.Category)

	if err != nil {
		log.Print("Delete category: ", err)
		return
	}

	c.JSON(http.StatusOK, midas.ApiAnswer{Status: 0})
}
