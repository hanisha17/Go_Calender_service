package controllers

import (
	"calender-service/models"
	"calender-service/services"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	service services.EventServiceInterface 
}

func NewEventController(service services.EventServiceInterface) *EventController {
	return &EventController{service: service}
}

func (c *EventController) CreateEvent(ctx *gin.Context) {
	log.Println("Received request to create an event")
	var event models.Event

	// Bind the JSON body to the event model
	if err := ctx.ShouldBindJSON(&event); err != nil {
		log.Printf("Error in binding JSON %v", err)
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
		log.Println("Validation failed: Start and end time required")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Start and end time are required"})
		return
	}

	if event.EndTime.Before(event.StartTime) {
		log.Println("Validation failed: End time cannot be before start time")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "End time cannot be before start time."})
		return
	}

	if err := c.service.CreateEvent(&event); err != nil {
		log.Printf("Error creating event %v", err)
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	log.Println("Event created successfully")
	ctx.JSON(http.StatusCreated, event)
}

// GetAllEvents handles GET requests to retrieve all events.
func (c *EventController) GetAllEvents(ctx *gin.Context) {
	events, err := c.service.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve events"})
		return
	}
	log.Println("Successfully retrieved all events")
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
		log.Printf("Error in binding JSON %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.UpdateEvent(uint(eventID), &updatedEvent); err != nil {
		log.Printf("Error while updating event %v", err)
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	log.Println("Successfully updated event")
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
		log.Println("Invalid start time format")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start time format"})
		return
	}
	end, err := time.Parse(time.RFC3339, endParam)
	if err != nil {
		log.Printf("Invalid end time format %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end time format"})
		return
	}

	// Convert userID from string to uint
	var uid uint
	if _, err := fmt.Sscan(userID, &uid); err != nil {
		log.Printf("Invalid user ID %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	events, err := c.service.GetEventsByUserAndDateRange(uid, start, end)
	if err != nil {
		log.Printf("Could not retrieve events %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve events"})
		return
	}
	log.Println("Retrieved events successfully")
	ctx.JSON(http.StatusOK, events)
}
