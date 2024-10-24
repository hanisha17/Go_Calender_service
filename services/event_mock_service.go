package services

import (
	"calender-service/models"
	"time"

	"github.com/stretchr/testify/mock"
)

type EventServiceInterface interface {
	CreateEvent(event *models.Event) error
	GetAllEvents() ([]models.Event, error)
	UpdateEvent(eventID uint, updatedEvent *models.Event) error
	GetEventsByUserAndDateRange(userID uint, start, end time.Time) ([]models.Event, error)
}
type MockEventService struct {
	mock.Mock
}

func (m *MockEventService) CreateEvent(event *models.Event) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *MockEventService) GetAllEvents() ([]models.Event, error) {
	args := m.Called()
	return args.Get(0).([]models.Event), args.Error(1)
}

func (m *MockEventService) UpdateEvent(eventID uint, event *models.Event) error {
	args := m.Called(eventID, event)
	return args.Error(0)
}

func (m *MockEventService) GetEventsByUserAndDateRange(userID uint, start, end time.Time) ([]models.Event, error) {
	args := m.Called(userID, start, end)
	return args.Get(0).([]models.Event), args.Error(1)
}


