package controllers

import (
	"calender-service/models"
	"calender-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	service *services.RoomService
}

func NewRoomController (service *services.RoomService) *RoomController {
	return &RoomController{service}
}

func (c *RoomController) CreateRoom(ctx *gin.Context) {
	var room models.Room
	if err := ctx.ShouldBindJSON(&room); err !=nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateRoom(&room); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to create Room"})
		return
	}

	ctx.JSON(http.StatusCreated, room)
}

func (c *RoomController) GetAllRooms(ctx *gin.Context){
	rooms, err := c.service.GetAll()
	if err !=nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":"Could not retrive rooms" })
		return
	}
	ctx.JSON(http.StatusOK,rooms)
}