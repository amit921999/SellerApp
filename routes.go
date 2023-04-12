package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orderManagement/models"
	"orderManagement/service"
)

func StartServer() {
	r := gin.Default()
	r.GET("/order", func(c *gin.Context) {
		var form models.FilterOrd
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		order := service.GetOrder(form)
		c.JSON(http.StatusOK, order)
	})

	r.POST("/order", func(c *gin.Context) {
		createOrder := models.Order{}
		err := c.ShouldBind(&createOrder)
		if err != nil {
			return
		}
		order := service.CreateOrder(createOrder)
		c.JSON(http.StatusOK, order)

	})

	r.PATCH("/order/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updateOrder models.UpdateOrder
		err := c.ShouldBind(&updateOrder)
		if err != nil {
			return
		}
		order := service.UpdateOrder(updateOrder, id)
		c.JSON(http.StatusOK, order)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
