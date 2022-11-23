package router

import (
	modelCtrl "assignment2/controller"

	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	router := gin.Default()
	router.POST("/api/v1/orders", modelCtrl.InsertOrders)
	router.GET("/api/v1/orders/", modelCtrl.ShowOrders)
	router.PUT("/api/v1/orders/:OrderID", modelCtrl.UpdateOrder)
	router.DELETE("/api/v1/orders/:OrderID", modelCtrl.DeleteOrder)
	return router
}
