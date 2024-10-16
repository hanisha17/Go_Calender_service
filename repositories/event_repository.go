package repositories

import (
	"calender-service/config"
	"calender-service/models"
	"time"
)

type EventRepository struct{}

func (r *EventRepository) Create(event *models.Event) error {
	return config.GetDB().Create(event).Error
}

func (r *EventRepository) Update(event *models.Event) error {
	return config.GetDB().Save(event).Error
}

func (r *EventRepository) GetByUserAnddateRange(userID uint, start, end time.Time) ([]models.Event, error) {
	var events []models.Event
	err := config.GetDB().
		Where("user_id = ? AND ((start_time BETWEEN ? AND ?) OR (end_time BETWEEN ? AND ?))", userID, start, end, start, end).
		Find(&events).Error
	return events, err
}

func (r *EventRepository) IsConflict(userID uint,start,end time.Time) bool {
	var count int64
	config.GetDB().Model(&models.Event{}).Where(&models.Event{}).
	Where("user_id = ? AND ((start_time < ? AND end_time > ?) OR (start_time < ? AND end_time > ?))", userID, end, start, end, start).
	Count(&count)
	return count > 0
}

func (r *EventRepository) GetAll() ([]models.Event, error) {
	var events []models.Event
	err := config.GetDB().Find(&events).Error
	return events, err
}
func (r *EventRepository) GetByID(id uint) (*models.Event, error) {
	var event models.Event
	if err := config.GetDB().First(&event, id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}
