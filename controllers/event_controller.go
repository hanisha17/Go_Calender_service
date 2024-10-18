package controllers

import (
	"calender-service/models"
	"calender-service/services"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	services *services.EventService
}

func NewEventController(service *services.EventService) *EventController {
	return &EventController{service}
}


func (c *EventController) CreateEvent(ctx *gin.Context) {
    var event models.Event

    // Bind the JSON body to the event model
    if err := ctx.ShouldBindJSON(&event); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Parse start_time and end_time from query parameters (string converted)
    startParam := ctx.Query("start_time")
    endParam := ctx.Query("end_time")

    // Parse the start and end time from query parameters
    if startParam != "" {
        event.StartTime, _ = time.Parse(time.RFC3339, startParam)
    }
    if endParam != "" {
        event.EndTime, _ = time.Parse(time.RFC3339, endParam)
    }


    // Validate StartTime and EndTime
    if event.StartTime.IsZero() || event.EndTime.IsZero() {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Start and end time are required"})
        return
    }

    if event.EndTime.Before(event.StartTime) {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "End time cannot be before start time."})
        return
    }

    if err := c.services.CreateEvent(&event); err != nil {
        ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, event)
}



// GetAllEvents handles GET requests to retrieve all events.
func (c *EventController) GetAllEvents(ctx *gin.Context) {
	events, err := c.services.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve events"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func (c *EventController) UpdateEvent(ctx *gin.Context) {
	eventIDStr := ctx.Param("id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var updatedEvent models.Event
	if err := ctx.ShouldBindJSON(&updatedEvent); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.services.UpdateEvent(uint(eventID), &updatedEvent); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedEvent)
}

func (c *EventController) GetEventsByUserAndDateRange(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	// Parse query parameters for start and end date
	startParam := ctx.Query("start_time")
	endParam := ctx.Query("end_time")

	// Parse the start and end time from the query parameters
	start, err := time.Parse(time.RFC3339, startParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start time format"})
		return
	}
	end, err := time.Parse(time.RFC3339, endParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end time format"})
		return
	}

	// Convert userID from string to uint
	var uid uint
	if _, err := fmt.Sscan(userID, &uid); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	events, err := c.services.GetEventsByUserAndDateRange(uid, start, end)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve events"})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

