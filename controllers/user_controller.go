package controllers

import (
	"calender-service/models"
	"calender-service/services"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.JSON(http.StatusCreated, user)
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

	// Retrieve events from the service
	events, err := c.services.GetEventsByUserAndDateRange(uid, start, end)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve events"})
		return
	}

	// Return the events as a JSON response
	ctx.JSON(http.StatusOK, events)
}
