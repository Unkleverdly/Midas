package handler

import (
	"midas/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()
	router.Use()
	auth := router.Group("auth/")
	auth.POST("/signIn", h.signIn)
	auth.POST("/signUp", h.signUp)
	auth.OPTIONS("/signUp", AddHeaders)
	auth.OPTIONS("/signIn", AddHeaders)

	user := router.Group("user/")
	user.POST("/newCategory", h.NewCategory)
	user.GET("/getCategories", h.getCategories)
	user.OPTIONS("/newCategory", AddHeaders)
	user.OPTIONS("/getCategories", AddHeaders)

	return router
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

// func AddHeaders(c *gin.Context) {
// 	c.Header("Access-Control-Allow-Origin", "*")
// 	c.Header("Access-Control-Allow-Headers", "*")
// }

func MiddleWare(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Next()
}
