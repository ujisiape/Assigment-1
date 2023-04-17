package controllers

import (
	"assignment-2/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertOrders(c *gin.Context) {
	var newOrder models.Orders
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newOrder = Insert(newOrder)
	c.JSON(http.StatusCreated, gin.H{
		"data":    newOrder,
		"message": "Data sucessfully created",
		"status":  http.StatusCreated,
	})
}

func ShowOrders(c *gin.Context) {
	orders := Show()
	c.JSON(http.StatusOK, gin.H{
		"data":    orders,
		"message": "Orders list fetched sucessfully",
		"status":  fmt.Sprintf("%d", http.StatusOK),
	})
}

func DeleteOrder(c *gin.Context) {
	OrderID := c.Param("OrderID")
	convertOrderID, err := strconv.Atoi(OrderID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":   nil,
			"status": fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}
	onDeleteID(uint(convertOrderID))
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Order with ID %v Has been sucessfully deleted", OrderID),
		"status":  fmt.Sprintf("%d", http.StatusOK),
	})
}

func UpdateOrder(c *gin.Context) {
	var updatedOrder models.Orders
	OrderID := c.Param("OrderID")
	convertOrderID, err := strconv.Atoi(OrderID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":   nil,
			"status": fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	updatedOrder = onUpdateID(updatedOrder, uint(convertOrderID))
	c.JSON(http.StatusOK, gin.H{
		"data":    updatedOrder,
		"message": fmt.Sprintf("Order with ID %v Has been sucessfully updated", OrderID),
		"status":  fmt.Sprintf("%d", http.StatusOK),
	})

}
