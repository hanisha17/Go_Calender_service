package repositories

import (
	"calender-service/config"
	"calender-service/models"
	"time"
)

type RoomRepository struct {
	CreateFunc      func(room *models.Room) error
	GetAllFunc      func() ([]models.Room, error)
}

func (r *RoomRepository) Create(room *models.Room) error {
	if r.CreateFunc != nil {
		return r.CreateFunc(room) 
	}
	return config.GetDB().Create(room).Error 
}

func (r *RoomRepository) GetAll() ([]models.Room, error) {
	if r.GetAllFunc != nil {
		return r.GetAllFunc() 
	}
	var rooms []models.Room
	err := config.GetDB().Find(&rooms).Error
	return rooms, err 
}

func (r *RoomRepository) IsRoomAvailable(roomID uint, start, end time.Time) bool {
	var count int64
	config.GetDB().Model(&models.Event{}).
		Where("room_id = ? AND ((start_time < ? AND end_time > ?) OR (start_time <? AND end_time > ?))", roomID, end, start, end, start).
		Count(&count)
	return count == 0
}
