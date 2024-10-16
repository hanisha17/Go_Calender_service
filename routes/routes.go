package routes

import (
	"calender-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController, eventController *controllers.EventController, roomController *controllers.RoomController) *gin.Engine {
	router := gin.Default()

	// User routes
	router.POST("/users", userController.CreateUser)
	router.GET("/users",userController.GetAllUsers)

	// Event routes
	router.POST("/events", eventController.CreateEvent)
	router.GET("/events", eventController.GetAllEvents)
	router.GET("/events/user/:user_id", eventController.GetEventsByUserAndDateRange)

	// Room routes
	router.POST("/rooms", roomController.CreateRoom)
	router.GET("/rooms",roomController.GetAllRooms)

	return router
}
