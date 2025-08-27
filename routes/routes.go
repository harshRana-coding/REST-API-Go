package routes

import (
	"rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/event",createEvent)
	authenticated.PUT("/event/:id", updateEvent)
	authenticated.DELETE("/event/:id", deleteEvent)
	authenticated.POST("/event/:id/register", registerForEvent)
	authenticated.DELETE("/event/:id/register", unregisterFromEvent)

	server.POST("/signup", signUp)
	server.POST("/login", login)
}
