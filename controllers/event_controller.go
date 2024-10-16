package controllers

import (
	"calender-service/models"
	"calender-service/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	services *services.EventService
}

func NewEventController(service *services.EventService) *EventController {
	return &EventController{service}
}

// func (c *EventController) CreateEvent(ctx *gin.Context) {
//     var event models.Event

//     // Bind the JSON body to the event model
//     if err := ctx.ShouldBindJSON(&event); err != nil {
//         ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }
// 	fmt.Println(event.StartTime)

//     // Validate StartTime and EndTime
//     if event.StartTime.IsZero() || event.EndTime.IsZero() {
//         ctx.JSON(http.StatusBadRequest, gin.H{"error": "Start and end time are required"})
//         return
//     }

//     // Call the service to create the event
//     if err := c.services.CreateEvent(&event); err != nil {
//         ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
//         return
//     }

//     // Return the created event as a response
//     ctx.JSON(http.StatusOK, event)
// }


func (c *EventController) CreateEvent(ctx *gin.Context) {
    var event models.Event

    // Bind the JSON body to the event model
    if err := ctx.ShouldBindJSON(&event); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Parse start_time and end_time from query parameters
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

    // Call the service to create the event
    if err := c.services.CreateEvent(&event); err != nil {
        ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        return
    }

    // Return the created event as a response
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

