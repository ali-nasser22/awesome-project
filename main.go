package main

import (
	"awesomeProject/db"
	"awesomeProject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	r := gin.Default()

	r.GET("/events", func(c *gin.Context) {
		events, err := models.GetAllEvents()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"events": events,
		})
	})

	r.POST("/events", func(c *gin.Context) {
		var event models.Event
		err := c.ShouldBindJSON(&event)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		createdStatus, err := event.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(createdStatus, gin.H{
			"message": "event created successfully",
		})
	})

	err := r.Run(":8080") // localhost:8080
	if err != nil {
		return
	}
}
