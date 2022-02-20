package routes

import (
	"github.com/ericklima-ca/go-gin-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func Handler() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api")
	{
		v1.GET("/orders", controllers.ListAllOrders)
		v1.GET("/orders/:ID", controllers.GetOrderByID)
		v1.POST("/orders", controllers.CreateOrder)
		v1.DELETE("/orders/:ID", controllers.DeleteOrderByID)
	}
	
	return r
}