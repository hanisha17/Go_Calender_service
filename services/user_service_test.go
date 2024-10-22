package services_test

import (
	"calender-service/models"
	"calender-service/repositories"
	"calender-service/services"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUserServie_CreateUser( t *testing.T){
	mockRepo:=&repositories.UserRepository{
		CreateFunc: func(user *models.User) error {
			return nil
		},
	}
	userService := services.NewUserService(mockRepo)
	user := &models.User{ID :1,Name:"User"}
	err := userService.CreateUser(user)
	assert.NoError(t,err)
}

func TestUserService_GetAll(t *testing.T){
	expectedUsers := []models.User{
		{ID:1,Name:"user"},
		{ID:2, Name:"user1"},
	}

	mockRepo := &repositories.UserRepository{
		GetAllFunc: func() ([]models.User, error) {
			return expectedUsers ,nil
		},
	}

	userService := services.NewUserService(mockRepo)

	user,err := userService.GetAllUsers()

	assert.Equal(t,expectedUsers,user)
	assert.NoError(t, err)
}

func TestUserService_GetUserById_UserFound(t *testing.T) {
	mockRepo := &repositories.UserRepository{
		GetAllId: func(id uint) (*models.User, error) {
			return &models.User{ID: 1, Name: "John Doe"}, nil
		},
	}
	userService := services.NewUserService(mockRepo)

	user, err := userService.GetUserById(1)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
	assert.Equal(t, "John Doe", user.Name)
}

func TestUserService_GetUserById_UserNotFound(t *testing.T) {
	mockRepo := &repositories.UserRepository{
		GetAllId: func(id uint) (*models.User, error) {
			return nil, gorm.ErrRecordNotFound
		},
	}
	userService := services.NewUserService(mockRepo)

	user, err := userService.GetUserById(2)

	assert.Error(t, err)
	assert.EqualError(t, err, "user not found")
	assert.Nil(t, user)
}

func TestUserService_GetUserById_OtherError(t *testing.T) {
	mockRepo := &repositories.UserRepository{
		GetAllId: func(id uint) (*models.User, error) {
			return nil, errors.New("database error")
		},
	}
	userService := services.NewUserService(mockRepo)

	user, err := userService.GetUserById(3)

	assert.Error(t, err)
	assert.Nil(t, user)
}