package models_test

import (
	"calender-service/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEventInitialization(t *testing.T) {
	startTime := time.Date(2024, 10, 12, 10, 0, 0, 0, time.UTC)
	endTime := time.Date(2024, 10, 12, 12, 0, 0, 0, time.UTC)

	event := models.Event{
		ID:        1,
		UserID:    1,
		Name:      "Test Event",
		StartTime: startTime,
		EndTime:   endTime,
		RoomID:    nil,
	}

	assert.Equal(t, uint(1), event.ID)
	assert.Equal(t, uint(1), event.UserID)
	assert.Equal(t, "Test Event", event.Name)
	assert.Equal(t, startTime, event.StartTime)
	assert.Equal(t, endTime, event.EndTime)
	assert.Nil(t, event.RoomID)
}
