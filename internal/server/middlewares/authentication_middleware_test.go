package middlewares

import (
	"RestApiBackend/infrastructure"
	"RestApiBackend/internal/features/users/entities"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FetchUserById(ctx context.Context, id string) (*entities.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) FetchUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) CreateNewUser(ctx context.Context, email string, firstName string, lastName string, password string) (*entities.User, error) {
	args := m.Called(ctx, email, firstName, lastName, password)
	return args.Get(0).(*entities.User), args.Error(1)
}

func TestAuthMiddleware(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Create the mock user repository
	mockUserRepo := new(MockUserRepository)

	// Mock the application
	app := &infrastructure.Application{
		Env: &infrastructure.Env{
			TokenSignerKey: "test_secret_key",
		},
	}

	// Test cases
	tests := []struct {
		name           string
		setupRequest   func(req *http.Request)
		tokenValidator TokenValidatorFunction
		expectedStatus int
		expectedBody   string
		mockSetup      func()
	}{
		{
			name: "Missing Authorization Header",
			setupRequest: func(req *http.Request) {
				// No Authorization header
			},
			tokenValidator: func(app *infrastructure.Application, accessToken string) (bool, *string, error) {
				return false, nil, nil
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"Authorization header is required"}`,
			mockSetup:      func() {},
		},
		{
			name: "Invalid Authorization Header Format",
			setupRequest: func(req *http.Request) {
				req.Header.Set("Authorization", "InvalidHeader")
			},
			tokenValidator: func(app *infrastructure.Application, accessToken string) (bool, *string, error) {
				return false, nil, nil
			},
			expectedStatus: http.StatusForbidden,
			expectedBody:   "",
			mockSetup:      func() {},
		},
		{
			name: "Invalid Token",
			setupRequest: func(req *http.Request) {
				req.Header.Set("Authorization", "Bearer invalid_token")
			},
			tokenValidator: func(app *infrastructure.Application, accessToken string) (bool, *string, error) {
				return false, nil, nil
			},
			expectedStatus: http.StatusForbidden,
			expectedBody:   "",
			mockSetup:      func() {},
		},
		{
			name: "Valid Token but User Fetch Fails",
			setupRequest: func(req *http.Request) {
				req.Header.Set("Authorization", "Bearer valid_token")
			},
			tokenValidator: func(app *infrastructure.Application, accessToken string) (bool, *string, error) {
				userID := "123"
				return true, &userID, nil
			},
			expectedStatus: http.StatusForbidden,
			expectedBody:   "",
			mockSetup: func() {
				user := &entities.User{}
				mockUserRepo.On("FetchUserById", mock.Anything, "123").Return(user, errors.New("user not found"))
			},
		},
		{
			name: "Valid Token and User Found",
			setupRequest: func(req *http.Request) {
				req.Header.Set("Authorization", "Bearer valid_token")
			},
			tokenValidator: func(app *infrastructure.Application, accessToken string) (bool, *string, error) {
				userID := "123"
				return true, &userID, nil
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "",
			mockSetup: func() {
				expectedUser := &entities.User{ID: uuid.New(), FirstName: "John", LastName: "Doe"}
				mockUserRepo.On("FetchUserById", mock.Anything, "123").Return(expectedUser, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup the middleware
			middleware := NewAuthBearerToken(mockUserRepo, app, tt.tokenValidator)

			// Create a gin router with the middleware
			r := gin.New()
			r.Use(middleware)
			r.GET("/test", func(c *gin.Context) {
				user, exists := c.Get("user")
				if exists {
					c.JSON(http.StatusOK, user)
				} else {
					c.Status(http.StatusInternalServerError)
				}
			})

			// Create a request to test the middleware
			req, err := http.NewRequest(http.MethodGet, "/test", nil)
			assert.NoError(t, err)

			// Setup request headers
			if tt.setupRequest != nil {
				tt.setupRequest(req)
			}

			// Setup mock expectations
			tt.mockSetup()

			// Create a response recorder to capture the response
			rr := httptest.NewRecorder()

			// Perform the request
			r.ServeHTTP(rr, req)

			// Check the status code and response body
			assert.Equal(t, tt.expectedStatus, rr.Code)
			if tt.expectedBody != "" {
				assert.JSONEq(t, tt.expectedBody, rr.Body.String())
			}

			// Assert that the mock was called as expected
			mockUserRepo.AssertExpectations(t)
		})
	}
}
