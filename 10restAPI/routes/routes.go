package routes

import (
	"example.com/example/10restAPI/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	//Separate approach with middleware
	// server.POST("/events", middlewares.Authenticate, createEvent)

	// or grouped
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup/", signup)
	server.POST("/login/", login)

}
