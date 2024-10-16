package services

import (
	"calender-service/models"
	"calender-service/repositories"
	"errors"
	"time"
)

type EventService struct {
	repo     *repositories.EventRepository
	roomRepo *repositories.RoomRepository
}

func NewEventService(repo *repositories.EventRepository, roomRepo *repositories.RoomRepository) *EventService {
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
