package controllers

import (
	"net/http"

	"github.com/ericklima-ca/go-gin-rest-api/database"
	"github.com/ericklima-ca/go-gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func ListAllOrders(c *gin.Context) {
	var orders []models.Order
	database.DB.Preload("Documents").Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func GetOrderByID(c *gin.Context) {
	var order models.Order
	id := c.Param("ID")
	database.DB.Preload("Documents").Find(&order, id)
	if order.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order not found",
		})
		return
	}
	c.JSON(http.StatusOK, order)
}

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&order)
	c.JSON(http.StatusCreated, order)
}

func DeleteOrderByID(c *gin.Context) {
	id := c.Param("ID")
	var order models.Order
	database.DB.Delete(&order, id)
	

	c.JSON(http.StatusNoContent, gin.H{})
}