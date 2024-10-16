package main

import (
	"calender-service/config"
	"calender-service/controllers"
	"calender-service/repositories"
	"calender-service/routes"
	"calender-service/services"
	"fmt"

)


func main(){
	config.InitDB();
	fmt.Println("hello")
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() 

	userRepo := &repositories.UserRepository{}
	eventRepo := &repositories.EventRepository{}
	roomRepo := &repositories.RoomRepository{}

	// Initialize services
	userService := services.NewUserService(userRepo)
	eventService := services.NewEventService(eventRepo, roomRepo)
	roomService := services.NewRoomService(roomRepo)

	// Initialize controllers
	userController := controllers.NewUserController(userService)
	eventController := controllers.NewEventController(eventService)
	roomController := controllers.NewRoomController(roomService)

	// Setup routes and start the server
	r := routes.SetupRouter(userController, eventController, roomController)
	r.Run() // Start the server
}