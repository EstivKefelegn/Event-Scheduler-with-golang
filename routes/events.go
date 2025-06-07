package routes

import (
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

	events.ID = 1
	events.UserID = 1

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

	_, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch the values from the db, try agin later"})
	}

}
