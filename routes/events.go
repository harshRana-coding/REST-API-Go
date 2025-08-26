package routes

import (
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}

func getEvent(context *gin.Context) {
	id := context.Param("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventById(intId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
	id := context.Param("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	_, err = models.GetEventById(intId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Could not retrieve event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEvent.ID = intId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated",
		"event":   updatedEvent,
	})

}

func deleteEvent(context *gin.Context) {
	id := context.Param("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	// check if the event exists
	_, err = models.GetEventById(intId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Could not retrieve event"})
		return
	}

	err = models.DeleteEvent(intId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted",
	})
}
