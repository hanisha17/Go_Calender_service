package services

import (
	"calender-service/models"
	"calender-service/repositories"
	"errors"
	"time"
)

type EventService struct {
	repo     repositories.EventRepositoryInterface
	roomRepo repositories.RoomRepositoryInterface
}

func NewEventService(repo repositories.EventRepositoryInterface, roomRepo repositories.RoomRepositoryInterface) *EventService {
	return &EventService{
		repo,
		roomRepo}
}

func (s *EventService) CreateEvent(event *models.Event) error {
	// Check for event conflicts
	if s.repo.IsConflict(event.UserID, event.StartTime, event.EndTime) {
		return errors.New("event time conflicts with another event")
	}

	// Check for room availability (optional)
	if event.RoomID != nil && !s.roomRepo.IsRoomAvailable(*event.RoomID, event.StartTime, event.EndTime) {
		return errors.New("room is already booked during this time")
	}

	return s.repo.Create(event)
}

// GetEventsByUserAndDateRange retrieves events for a user within a specified date range.
func (s *EventService) GetEventsByUserAndDateRange(userID uint, start, end time.Time) ([]models.Event, error) {
	return s.repo.GetByUserAnddateRange(userID, start, end)
}

// GetAllEvents retrieves all events.
func (s *EventService) GetAllEvents() ([]models.Event, error) {
	return s.repo.GetAll() 
}

//UpdateEvent checks for conflict and update the event
func (s *EventService) UpdateEvent(eventID uint, updatedEvent *models.Event) error {
	event, err := s.repo.GetByID(eventID)
	if err !=nil {
		return errors.New("event not found")
	}

	event.Name = updatedEvent.Name
	event.StartTime = updatedEvent.StartTime
	event.EndTime = updatedEvent.EndTime
	event.RoomID = updatedEvent.RoomID

	if s.repo.IsConflict(event.UserID, event.StartTime, event.EndTime) {
		return errors.New("event time conflicts with another event")
	}

	// Check for room availability (if RoomID is present)
	if event.RoomID != nil && !s.roomRepo.IsRoomAvailable(*event.RoomID, event.StartTime, event.EndTime) {
		return errors.New("room is already booked during this time")
	}
	return s.repo.Update(event)
}




