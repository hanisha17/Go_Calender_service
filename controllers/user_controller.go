package controllers

import (
	"calender-service/models"
	"calender-service/services"
	"log"
	"net/http"

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
		log.Printf("%v",err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.CreateUser(&user); err != nil {
		log.Printf("%v",err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	log.Println("succesfully created user")
	ctx.JSON(http.StatusCreated, user)
}


func (c *UserController) GetAllUsers (ctx *gin.Context) {
	users, err := c.service.GetAllUsers()
	if err != nil {
		log.Printf("%v",err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Counld not retrive users"})
		return
	}
	log.Println("Users created")
	ctx.JSON(http.StatusOK, users)
}
