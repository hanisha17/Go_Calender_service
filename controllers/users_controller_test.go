package controllers

import (
	"bytes"
	"calender-service/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock the UserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) GetAllUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	router := gin.Default()
	mockService := new(MockUserService)
	controller := NewUserController(mockService)
	router.POST("/users", controller.CreateUser)

	tests := []struct {
		name           string
		input          models.User
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Successful User Creation",
			input: models.User{
				ID:   1,
				Name: "John Doe",
			},
			mockError:      nil,
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"id":1,"name":"John Doe"}`,
		},
		{
			name:           "Failed User Creation due to Service Error",
			input:          models.User{Name: "John Doe"},
			mockError:      assert.AnError,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"Failed to create user"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockService.ExpectedCalls = nil
			mockService.On("CreateUser", &tc.input).Return(tc.mockError)

			// Convert input to JSON
			body, _ := json.Marshal(tc.input)
			req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.JSONEq(t, tc.expectedBody, w.Body.String())
			mockService.AssertExpectations(t)
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	router := gin.Default()
	mockService := new(MockUserService)
	controller := NewUserController(mockService)
	router.GET("/users", controller.GetAllUsers)

	tests := []struct {
		name           string
		mockUsers      []models.User
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Successful Retrieval of Users",
			mockUsers: []models.User{
				{ID: 1, Name: "John Doe"},
				{ID: 2, Name: "Jane Doe"},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"id":1,"name":"John Doe"},{"id":2,"name":"Jane Doe"}]`,
		},
		{
			name:           "Error Retrieving Users",
			mockUsers:      nil,
			mockError:      assert.AnError,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"Counld not retrive users"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockService.ExpectedCalls = nil
			mockService.On("GetAllUsers").Return(tc.mockUsers, tc.mockError)

			req, _ := http.NewRequest(http.MethodGet, "/users", nil)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.JSONEq(t, tc.expectedBody, w.Body.String())
			mockService.AssertExpectations(t)
		})
	}
}
