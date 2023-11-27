package handler

import (
	"midas/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()
	router.Use(h.AddHeaders)

	auth := router.Group("auth/")
	auth.POST("/signIn", h.signIn)
	auth.POST("/signUp", h.signUp)

	user := router.Group("user/")
	// user.GET("/getUser", h.getUser)
	user.POST("/getMainData", h.getMainData)

	user.POST("/addCategory", h.addCategory)
	user.POST("/deleteCategory", h.deleteCategory)
	user.POST("/getCategories", h.getCategories)

	user.POST("/makeTransaction", h.makeTransaction)
	// user.POST("/getTranscation", h.)

	return router
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) AddHeaders(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")

	if c.Request.Method == "OPTIONS" {
		c.Status(http.StatusOK)
		return
	}

	c.Next()
}
