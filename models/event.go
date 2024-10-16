package models

import "time"

type Event struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	StartTime time.Time `gorm:"not null" json:"start_time"`
	EndTime   time.Time `gorm:"not null" json:"end_time"`
	// User      User      `gorm:"foreignKey:UserID" json:"user"`

	// Optional room assignment
	RoomID *uint `gorm:"null" json:"room_id"`
	Room   *Room `gorm:"foreignKey:RoomID" json:"room"`
}
