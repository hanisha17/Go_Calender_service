package controllers

import (
	"bytes"
	"calender-service/models"
	"calender-service/services"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateEvent(t *testing.T) {
	router := gin.Default()
	mockService := new(services.MockEventService)
	controller := NewEventController(mockService)
	router.POST("/events", controller.CreateEvent)

	tests := []struct {
		name           string
		eventData      []byte
		startParam     string
		endParam       string
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Successful Event Creation",
			eventData: []byte(`{
				"name": "Test Event",
				"user_id": 1,
				"start_time": "2024-10-24T15:13:05Z",
				"end_time": "2024-10-24T17:13:05Z"
			}`),
			startParam:     "2024-10-24T15:13:05Z",
			endParam:       "2024-10-24T17:13:05Z",
			mockError:      nil,
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "Invalid JSON",
			eventData:      []byte(``),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"EOF"}`,
		},
		{
			name: "Missing Start and End Time",
			eventData: []byte(`{
				"name": "Test Event",
				"user_id": 1
			}`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Start and end time are required"}`,
		},
		{
			name: "End Time Before Start Time",
			eventData: []byte(`{
				"name": "Invalid Event",
				"user_id": 1,
				"start_time": "2024-10-24T17:13:05Z",
				"end_time": "2024-10-24T15:13:05Z"
			}`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"End time cannot be before start time."}`,
		},
		{
			name: "Event Conflict",
			eventData: []byte(`{
				"name": "Conflicting Event",
				"user_id": 1,
				"start_time": "2024-10-24T15:13:05Z",
				"end_time": "2024-10-24T17:13:05Z"
			}`),
			startParam:     "2024-10-24T15:13:05Z",
			endParam:       "2024-10-24T17:13:05Z",
			mockError:      assert.AnError,
			expectedStatus: http.StatusConflict,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.mockError != nil {
				mockService.On("CreateEvent", mock.AnythingOfType("*models.Event")).Return(tc.mockError)
			} else {
				mockService.On("CreateEvent", mock.AnythingOfType("*models.Event")).Return(nil)
			}

			req, _ := http.NewRequest(http.MethodPost, "/events", bytes.NewBuffer(tc.eventData))
			req.Header.Set("Content-Type", "application/json")

			if tc.startParam != "" || tc.endParam != "" {
				q := req.URL.Query()
				if tc.startParam != "" {
					q.Add("start_time", tc.startParam)
				}
				if tc.endParam != "" {
					q.Add("end_time", tc.endParam)
				}
				req.URL.RawQuery = q.Encode()
			}

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if tc.expectedBody != "" {
				assert.Contains(t, w.Body.String(), tc.expectedBody)
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestUpdateEvent(t *testing.T) {
	router := gin.Default()
	mockService := new(services.MockEventService)
	controller := NewEventController(mockService)
	router.PUT("/events/:id", controller.UpdateEvent)

	tests := []struct {
		name           string
		eventID        string
		eventData      models.Event
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Successful Update",
			eventID:        "1",
			eventData:      models.Event{Name: "Updated Event", StartTime: mockTime(), EndTime: mockTime()},
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid Event ID",
			eventID:        "invalid",
			eventData:      models.Event{Name: "Updated Event", StartTime: mockTime(), EndTime: mockTime()},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid event ID"}`,
		},
		{
			name:           "Event Update Conflict",
			eventID:        "2",
			eventData:      models.Event{Name: "Conflicting Event", StartTime: mockTime(), EndTime: mockTime()},
			mockError:      assert.AnError,
			expectedStatus: http.StatusConflict,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.mockError != nil {
				mockService.On("UpdateEvent", mock.AnythingOfType("uint"), mock.AnythingOfType("*models.Event")).Return(tc.mockError)
			} else {
				mockService.On("UpdateEvent", mock.AnythingOfType("uint"), mock.AnythingOfType("*models.Event")).Return(nil)
			}

			reqBody, _ := json.Marshal(tc.eventData)
			req, _ := http.NewRequest(http.MethodPut, "/events/"+tc.eventID, bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			if tc.expectedBody != "" {
				assert.Contains(t, w.Body.String(), tc.expectedBody)
			}

			mockService.AssertExpectations(t)
		})
	}
}

func mockTime() time.Time {
	return time.Now()
}

// func TestGetEventsByUserAndDateRange(t *testing.T) {
// 	// Initialize Gin router and mocks
// 	router := gin.Default()
// 	mockService := new(services.MockEventService)
// 	controller := NewEventController(mockService)
// 	router.GET("/events/user/:user_id", controller.GetEventsByUserAndDateRange)

// 	// Define test cases
// 	tests := []struct {
// 		name           string
// 		userID         string
// 		startTime      string
// 		endTime        string
// 		mockEvents     []models.Event
// 		mockError      error
// 		expectedStatus int
// 		expectedBody   string
// 	}{
// 		{
// 			name:           "Successful Retrieval",
// 			userID:         "1",
// 			startTime:      time.Now().Format(time.RFC3339),
// 			endTime:        time.Now().Add(2 * time.Hour).Format(time.RFC3339),
// 			mockEvents:     []models.Event{{Name: "Event1"}, {Name: "Event2"}},
// 			expectedStatus: http.StatusOK,
// 		},
// 		{
// 			name:           "Invalid Start Time Format",
// 			userID:         "1",
// 			startTime:      "invalid",
// 			endTime:        time.Now().Add(2 * time.Hour).Format(time.RFC3339),
// 			expectedStatus: http.StatusBadRequest,
// 			expectedBody:   `{"error":"Invalid start time format"}`,
// 		},
// 		{
// 			name:           "Service Error",
// 			userID:         "2",
// 			startTime:      time.Now().Format(time.RFC3339),
// 			endTime:        time.Now().Add(2 * time.Hour).Format(time.RFC3339),
// 			mockError:      assert.AnError,
// 			expectedStatus: http.StatusInternalServerError,
// 			expectedBody:   `{"error":"Could not retrieve events"}`,
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			// Mock service method
// 			if tc.mockError != nil {
// 				mockService.On("GetEventsByUserAndDateRange", mock.AnythingOfType("uint"), mock.Anything, mock.Anything).Return(nil, tc.mockError)
// 			} else {
// 				mockService.On("GetEventsByUserAndDateRange", mock.AnythingOfType("uint"), mock.Anything, mock.Anything).Return(tc.mockEvents, nil)
// 			}

// 			// Create the request and response
// 			req, _ := http.NewRequest(http.MethodGet, "/events/user/"+tc.userID+"?start_time="+tc.startTime+"&end_time="+tc.endTime, nil)
// 			w := httptest.NewRecorder()

// 			// Perform the request
// 			router.ServeHTTP(w, req)

// 			// Assert the response
// 			assert.Equal(t, tc.expectedStatus, w.Code)
// 			if tc.expectedBody != "" {
// 				assert.Contains(t, w.Body.String(), tc.expectedBody)
// 			}

// 			mockService.AssertExpectations(t)
// 		})
// 	}
// }

func TestGetAllEvents(t *testing.T) {
	router := gin.Default()
	mockService := new(services.MockEventService)
	controller := NewEventController(mockService)
	router.GET("/events", controller.GetAllEvents)

	tests := []struct {
		name           string
		mockEvents     []models.Event
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Successful Retrieval of Events",
			mockEvents: []models.Event{
				{ID: 1, UserID: 1, Name: "Event 1", StartTime: time.Now(), EndTime: time.Now().Add(1 * time.Hour), RoomID: nil},
				{ID: 2, UserID: 1, Name: "Event 2", StartTime: time.Now(), EndTime: time.Now().Add(1 * time.Hour), RoomID: nil},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"id":1,"user_id":1,"name":"Event 1","start_time":"<start_time>","end_time":"<end_time>","room_id":null},{"id":2,"user_id":1,"name":"Event 2","start_time":"<start_time>","end_time":"<end_time>","room_id":null}]`,
		},
		{
			name:           "Error Retrieving Events",
			mockEvents:     nil,
			mockError:      fmt.Errorf("database error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"Could not retrieve events"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockService.On("GetAllEvents").Return(tc.mockEvents, tc.mockError)

			req, _ := http.NewRequest(http.MethodGet, "/events", nil)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.expectedBody != "" {
				assert.JSONEq(t, tc.expectedBody, w.Body.String())
			}

			mockService.AssertExpectations(t)
		})
	}
}
