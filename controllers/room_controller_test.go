package controllers

import (
	"bytes"
	"calender-service/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRoomService is a mock implementation of the RoomServiceInterface
type MockRoomService struct {
	mock.Mock
}

func (m *MockRoomService) CreateRoom(room *models.Room) error {
	args := m.Called(room)
	return args.Error(0)
}

func (m *MockRoomService) GetAll() ([]models.Room, error) {
	args := m.Called()
	return args.Get(0).([]models.Room), args.Error(1)
}

func TestRoomController_CreateRoom(t *testing.T) {
	mockService := new(MockRoomService)
	controller := NewRoomController(mockService)

	router := gin.Default()
	router.POST("/rooms", controller.CreateRoom)

	t.Run("Successfully Create Room", func(t *testing.T) {
		mockRoom := models.Room{
			ID:   1,
			Name: "Room 101",
		}
		mockService.On("CreateRoom", &mockRoom).Return(nil)

		body, _ := json.Marshal(mockRoom)
		req, _ := http.NewRequest(http.MethodPost, "/rooms", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("Failed to Create Room", func(t *testing.T) {
		mockRoom := models.Room{}
		mockService.On("CreateRoom", &mockRoom)

		body, _ := json.Marshal(mockRoom)
		req, _ := http.NewRequest(http.MethodPost, "/rooms", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockService.AssertExpectations(t)
	})
}

func TestRoomController_GetAllRooms(t *testing.T) {

	router := gin.Default()
	mockService := new(MockRoomService)
	controller := NewRoomController(mockService)
	router.GET("/rooms", controller.GetAllRooms)

	t.Run("Successfully Retrieve All Rooms", func(t *testing.T) {
		mockRooms := []models.Room{
			{ID: 1, Name: "Room 101"},
			{ID: 2, Name: "Room 102"},
		}
		mockService.On("GetAll").Return(mockRooms, nil)

		req, _ := http.NewRequest(http.MethodGet, "/rooms", nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("Failed to Retrieve Rooms", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.On("GetAll").Return(nil, fmt.Errorf("database error"))

		req, _ := http.NewRequest(http.MethodGet, "/rooms", nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockService.AssertExpectations(t)
	})
}
