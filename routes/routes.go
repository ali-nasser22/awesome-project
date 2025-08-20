package routes

import (
	"awesomeProject/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	/*Events Routes*/

	r.GET("/events", getEvents)
	r.GET("/events/:id", getEvent)

	/* Authenticated Routes */
	authenticated := r.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", saveEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	/* Users Routes*/
	r.POST("/signup", saveUser)
	r.POST("/login", login)
}
