package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	/*Events Routes*/

	r.GET("/events", getEvents)
	r.GET("/events/:id", getEvent)
	r.POST("/events", saveEvent)
	r.PUT("/events/:id", updateEvent)
	r.DELETE("/events/:id", deleteEvent)
}
