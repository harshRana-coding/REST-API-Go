package routes

import (
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId := context.Param("id")
	intEventId, err := strconv.ParseInt(eventId, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	// check if the event exists
	event, err := models.GetEventById(intEventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Could not retrieve event"})
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register for event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully registered for event"})
}

func unregisterFromEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId := context.Param("id")
	intEventId, err := strconv.ParseInt(eventId, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	// check if the event exists
	event, err := models.GetEventById(intEventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Could not retrieve event"})
		return
	}
	err = event.Unregister(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not unregister from event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully unregistered from event"})
}
