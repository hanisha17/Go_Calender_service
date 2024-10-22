package services_test

import (
	"calender-service/models"
	"calender-service/repositories"
	"calender-service/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoomService_CreateRoom(t *testing.T) {
	mockRepo := &repositories.RoomRepository{
		CreateFunc: func(room *models.Room) error {
			return nil
		},
	}
	roomService := services.NewRoomService(mockRepo)

	room := &models.Room{ID: 1, Name: "Conference Room"}

	err := roomService.CreateRoom(room)

	assert.NoError(t, err)
}

func TestRoomService_GetAll(t *testing.T) {
	expectedRooms := []models.Room{
		{ID: 1, Name: "Conference Room"},
		{ID: 2, Name: "Meeting Room"},
	}

	mockRepo := &repositories.RoomRepository{
		GetAllFunc: func() ([]models.Room, error) {
			return expectedRooms, nil
		},
	}
	roomService := services.NewRoomService(mockRepo)

	rooms, err := roomService.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedRooms, rooms)
}

// func TestRoomService_GetAll_Error(t *testing.T) {
// 	mockRepo := &repositories.RoomRepository{
// 		GetAllFunc: func() ([]models.Room, error) {
// 			// Mock the behavior of GetAll to return an error
// 			return nil, errors.New("failed to fetch rooms")
// 		},
// 	}
// 	roomService := services.NewRoomService(mockRepo)

// 	// Test GetAll
// 	rooms, err := roomService.GetAll()

// 	// Assertions
// 	assert.Error(t, err)
// 	assert.Nil(t, rooms)
// }
