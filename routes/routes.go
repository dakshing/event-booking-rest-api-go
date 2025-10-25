package routes

import (
	"github.com/dakshing/event-booking-rest-api-go/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Protected Event routes
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// Event routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// User routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}
