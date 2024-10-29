package services

import (
	"calender-service/models"

	"github.com/stretchr/testify/mock"
)

type RoomServiceInterface interface {
	CreateRoom(room *models.Room) error
	GetAll() ([]models.Room, error)
}

type MockRoomService struct {
	mock.Mock
}

func (m *MockUserService) CreateRoom(room *models.Room) error {
	args := m.Called(room)
	return args.Error(0)
}

func (m *MockRoomService) GetAll() ([]models.Room, error) {
	args := m.Called()
	return args.Get(0).([]models.Room), args.Error(1)
}
