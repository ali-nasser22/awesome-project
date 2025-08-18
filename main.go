package main

import (
	"awesomeProject/db"
	"awesomeProject/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	r := gin.Default()

	routes.RegisterRoutes(r)

	err := r.Run(":8080") // localhost:8080
	if err != nil {
		return
	}
}
