package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"wallet-app/backend/models"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	// Clear in-memory stores before each test run
	models.Users = make(map[string]models.User)
	models.UserEmailIndex = make(map[string]string)

	t.Run("Successful Registration", func(t *testing.T) {
		payload := RegisterRequest{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password123",
		}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(RegisterHandler)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code, "Status code should be 201 Created")

		var resp RegisterResponse
		err := json.NewDecoder(rr.Body).Decode(&resp)
		assert.NoError(t, err, "Should decode response body without error")
		assert.Equal(t, "Test User", resp.Name)
		assert.Equal(t, "test@example.com", resp.Email)
		assert.NotEmpty(t, resp.UserID, "UserID should not be empty")

		// Verify user was stored (basic check)
		_, existsInIndex := models.UserEmailIndex["test@example.com"]
		assert.True(t, existsInIndex, "User email should be in the index")
	})

	t.Run("Email Already Exists", func(t *testing.T) {
		// First, register a user
		existingUser := RegisterRequest{Name: "Existing User", Email: "exists@example.com", Password: "password123"}
		body1, _ := json.Marshal(existingUser)
		req1, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(body1))
		rr1 := httptest.NewRecorder()
		RegisterHandler(rr1, req1) // Directly call, or use a mux if prefix stripping is an issue
		assert.Equal(t, http.StatusCreated, rr1.Code)


		// Attempt to register the same email again
		payload := RegisterRequest{Name: "Another User", Email: "exists@example.com", Password: "password456"}
		body2, _ := json.Marshal(payload)
		req2, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(body2))
		rr2 := httptest.NewRecorder()
		RegisterHandler(rr2, req2)

		assert.Equal(t, http.StatusConflict, rr2.Code, "Status code should be 409 Conflict")
	})

	t.Run("Invalid Input - Missing Fields", func(t *testing.T) {
		payload := RegisterRequest{Name: "Test User", Email: "", Password: "password123"} // Missing email
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		RegisterHandler(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code, "Status code for missing fields")
	})

	t.Run("Invalid Input - Password too short", func(t *testing.T) {
		payload := RegisterRequest{Name: "Test User", Email: "shortpass@example.com", Password: "pass"} // Password too short
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		RegisterHandler(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code, "Status code for short password")
	})
}
