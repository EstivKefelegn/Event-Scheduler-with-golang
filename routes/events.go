package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"Eventplanning.go/Api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch the values from the db, try agin later"})
	}

	

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't fetch the values from the db, try agin later"})
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch the values from the db, try agin later"})
	}
	context.JSON(http.StatusOK, event)
}

func createEvents(context *gin.Context) {

	var events models.Event
	err := context.ShouldBindJSON(&events)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// events.ID = 1
	userID := context.GetInt64("userID")
	events.UserID = userID

	err = events.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch the values from the db, try agin later"})
	}
	context.JSON(http.StatusCreated, gin.H{"Message": "Event created", "events": events})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't fetch the values from the db, try agin later"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch the values from the db, try agin later"})
	}

	userID := context.GetInt64("userID")

	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User is not authroized to update this event"})
		return
	}

	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "coundn't parse request data", "error": err.Error()})
		return
	}

	updateEvent.ID = eventId
	err = updateEvent.UpdateEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't update events"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	fmt.Println("Received event id:", context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't pare the event id"})
		return
	}

	event, err := models.GetEventById(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't fetch the event"})
		return
	}

	userID := context.GetInt64("userID")

	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User is not authroized to delete this event"})
		return
	}


	err = event.DeleteEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldnt fetch the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})

}
