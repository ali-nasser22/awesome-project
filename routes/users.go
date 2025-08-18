package routes

import (
	"awesomeProject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func saveUser(c *gin.Context) {

	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}
