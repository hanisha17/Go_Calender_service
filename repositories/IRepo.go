package repositories

import (
	"calender-service/models"
	"time"
)

type EventRepositoryInterface interface {
	Create(event *models.Event) error
	IsConflict(userID uint, start, end time.Time) bool
	GetByID(id uint) (*models.Event, error)
	GetAll() ([]models.Event, error)
	GetByUserAnddateRange(userID uint, start, end time.Time) ([]models.Event, error)
	Update(event *models.Event) error
}

type RoomRepositoryInterface interface {
	IsRoomAvailable(roomID uint, start, end time.Time) bool
}
