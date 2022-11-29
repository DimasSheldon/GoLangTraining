package router

import (
	"assignment3/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("view/*.html")

	router.GET("/", controller.GetStatus)

	return router
}
