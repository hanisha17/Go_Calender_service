package services_test

import (
	"calender-service/models"
	"calender-service/services"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEventRepository struct {
	mock.Mock
}

func (m *MockEventRepository) Create(event *models.Event) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *MockEventRepository) IsConflict(userID uint, start, end time.Time) bool {
	args := m.Called(userID, start, end)
	return args.Bool(0)
}

func (m *MockEventRepository) GetByID(id uint) (*models.Event, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Event), args.Error(1)
}

func (m *MockEventRepository) GetAll() ([]models.Event, error) {
	args := m.Called()
	return args.Get(0).([]models.Event), args.Error(1)
}

func (m *MockEventRepository) GetByUserAnddateRange(userID uint, start, end time.Time) ([]models.Event, error) {
	args := m.Called(userID, start, end)
	return args.Get(0).([]models.Event), args.Error(1)
}

func (m *MockEventRepository) Update(event *models.Event) error {
	args := m.Called(event)
	return args.Error(0)
}

type MockRoomRepository struct {
	mock.Mock
}

func (m *MockRoomRepository) IsRoomAvailable(roomID uint, start, end time.Time) bool {
	args := m.Called(roomID, start, end)
	return args.Bool(0)
}


func TestEventService_CreateEvent(t *testing.T) {
	mockEventRepo := new(MockEventRepository)
	mockRoomRepo := new(MockRoomRepository)
	eventService := services.NewEventService(mockEventRepo, mockRoomRepo)

	RoomId := uint(2)
	event := &models.Event{
		UserID:    1,
		RoomID: &RoomId,
		StartTime: time.Now(),
		EndTime:   time.Now().Add(1 * time.Hour),
	}

	// Mock the behavior
	mockEventRepo.On("IsConflict", event.UserID, event.StartTime, event.EndTime).Return(false)
	mockRoomRepo.On("IsRoomAvailable", *event.RoomID, event.StartTime, event.EndTime).Return(true)
	mockEventRepo.On("Create", event).Return(nil)

	// Call the service method
	err := eventService.CreateEvent(event)

	// Assertions
	assert.NoError(t, err)
	mockEventRepo.AssertExpectations(t)
	mockRoomRepo.AssertExpectations(t)
}

func TestGetEventsByUserAndDateRange(t *testing.T) {
    mockEventRepo := new(MockEventRepository)
	mockRoomRepo := new(MockRoomRepository)


	eventService := services.NewEventService(mockEventRepo,mockRoomRepo)
    userID := uint(1)
    start := time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2024, 10, 31, 23, 59, 59, 0, time.UTC)

    mockEvents := []models.Event{
        {
            ID:        1,
            UserID:    userID,
            Name:     "Event 1",
            StartTime: time.Date(2024, 10, 10, 10, 0, 0, 0, time.UTC),
            EndTime:   time.Date(2024, 10, 10, 12, 0, 0, 0, time.UTC),
        },
        {
            ID:        2,
            UserID:    userID,
            Name:     "Event 2",
            StartTime: time.Date(2024, 10, 20, 15, 0, 0, 0, time.UTC),
            EndTime:   time.Date(2024, 10, 20, 17, 0, 0, 0, time.UTC),
        },
    }

    mockEventRepo.On("GetByUserAnddateRange", userID, start, end).Return(mockEvents, nil)

    events, err := eventService.GetEventsByUserAndDateRange(userID, start, end)

    assert.NoError(t, err)
    assert.NotNil(t, events)
    assert.Equal(t, len(mockEvents), len(events))
    assert.Equal(t, mockEvents[0].Name, events[0].Name)
    assert.Equal(t, mockEvents[1].Name, events[1].Name)

    mockEventRepo.AssertCalled(t, "GetByUserAnddateRange", userID, start, end)
}


func TestGetAllEvents(t *testing.T) {
	  mockEventRepo := new(MockEventRepository)
	  mockRoomRepo := new(MockRoomRepository)
  
  
	  eventService := services.NewEventService(mockEventRepo,mockRoomRepo)

    mockEvents := []models.Event{
        {
            ID:        1,
            UserID:    1,
            Name:     "Event 1",
            StartTime: time.Date(2024, 10, 10, 10, 0, 0, 0, time.UTC),
            EndTime:   time.Date(2024, 10, 10, 12, 0, 0, 0, time.UTC),
        },
        {
            ID:        2,
            UserID:    2,
            Name:     "Event 2",
            StartTime: time.Date(2024, 10, 20, 15, 0, 0, 0, time.UTC),
            EndTime:   time.Date(2024, 10, 20, 17, 0, 0, 0, time.UTC),
        },
    }

    mockEventRepo.On("GetAll").Return(mockEvents, nil)

    events, err := eventService.GetAllEvents()

    assert.NoError(t, err)
    assert.NotNil(t, events)
    assert.Equal(t, len(mockEvents), len(events))
    assert.Equal(t, mockEvents[0].Name, events[0].Name)
    assert.Equal(t, mockEvents[1].Name, events[1].Name)

    mockEventRepo.AssertCalled(t, "GetAll")
}

func TestUpdateEvent(t *testing.T) {
	mockEventRepo := new(MockEventRepository)
	mockRoomRepo := new(MockRoomRepository)


	eventService := services.NewEventService(mockEventRepo,mockRoomRepo)

    eventID := uint(1)
    updatedEvent := &models.Event{
        Name:      "Updated Event",
        StartTime: time.Date(2024, 10, 20, 10, 0, 0, 0, time.UTC),
        EndTime:   time.Date(2024, 10, 20, 12, 0, 0, 0, time.UTC),
        RoomID:    func() *uint { r := uint(2); return &r }(),
    }

    existingEvent := &models.Event{
        ID:        eventID,
        UserID:    1,
        Name:      "Original Event",
        StartTime: time.Date(2024, 10, 20, 9, 0, 0, 0, time.UTC),
        EndTime:   time.Date(2024, 10, 20, 11, 0, 0, 0, time.UTC),
        RoomID:    nil,
    }

    t.Run("EventNotFound", func(t *testing.T) {
		mockEventRepo.ExpectedCalls = nil

		mockEventRepo.On("GetByID", eventID).Return((*models.Event)(nil), errors.New("event not found"))

		err := eventService.UpdateEvent(eventID, updatedEvent)
		assert.Error(t, err)
		assert.Equal(t, "event not found", err.Error())

		mockEventRepo.AssertExpectations(t)
	})

	t.Run("TimeConflict", func(t *testing.T) {
		mockEventRepo.ExpectedCalls = nil

		mockEventRepo.On("GetByID", eventID).Return(existingEvent, nil)
		mockEventRepo.On("IsConflict", existingEvent.UserID, updatedEvent.StartTime, updatedEvent.EndTime).Return(true)

		err := eventService.UpdateEvent(eventID, updatedEvent)
		assert.Error(t, err)
		assert.Equal(t, "event time conflicts with another event", err.Error())

		mockEventRepo.AssertExpectations(t)
	})

	t.Run("RoomNotAvailable", func(t *testing.T) {
		mockEventRepo.ExpectedCalls = nil
		mockRoomRepo.ExpectedCalls = nil

		mockEventRepo.On("GetByID", eventID).Return(existingEvent, nil)
		mockEventRepo.On("IsConflict", existingEvent.UserID, updatedEvent.StartTime, updatedEvent.EndTime).Return(false)
		mockRoomRepo.On("IsRoomAvailable", *updatedEvent.RoomID, updatedEvent.StartTime, updatedEvent.EndTime).Return(false)

		err := eventService.UpdateEvent(eventID, updatedEvent)
		assert.Error(t, err)
		assert.Equal(t, "room is already booked during this time", err.Error())

		mockEventRepo.AssertExpectations(t)
		mockRoomRepo.AssertExpectations(t)
	})

	t.Run("SuccessfulUpdate", func(t *testing.T) {
		mockEventRepo.ExpectedCalls = nil
		mockRoomRepo.ExpectedCalls = nil

		mockEventRepo.On("GetByID", eventID).Return(existingEvent, nil)
		mockEventRepo.On("IsConflict", existingEvent.UserID, updatedEvent.StartTime, updatedEvent.EndTime).Return(false)
		mockRoomRepo.On("IsRoomAvailable", *updatedEvent.RoomID, updatedEvent.StartTime, updatedEvent.EndTime).Return(true)

		mockEventRepo.On("Update", mock.MatchedBy(func(event *models.Event) bool {
			return event.Name == updatedEvent.Name &&
				event.StartTime == updatedEvent.StartTime &&
				event.EndTime == updatedEvent.EndTime &&
				event.RoomID != nil && *event.RoomID == *updatedEvent.RoomID
		})).Return(nil)

		err := eventService.UpdateEvent(eventID, updatedEvent)
		assert.NoError(t, err)

		mockEventRepo.AssertExpectations(t)
		mockRoomRepo.AssertExpectations(t)
	})
}