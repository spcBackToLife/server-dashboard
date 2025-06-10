package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"wallet-app/backend/models" // Adjust path if needed
	"github.com/stretchr/testify/assert"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Helper to register a user for login tests
func setupTestUserForLogin(email, password, name string) (models.User, error) {
	hashedPassword, err := models.HashPassword(password)
	if err != nil {
		return models.User{}, err
	}
	user := models.User{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}
	models.Users[user.ID] = user
	models.UserEmailIndex[user.Email] = user.ID
	return user, nil
}


func TestLoginHandler(t *testing.T) {
	// Reset in-memory stores and set a consistent JWT key for testing
	originalJwtKey := jwtKey
	jwtKey = []byte("test_secret_key_for_login_handler") // Consistent key for predictable tokens if needed, or use original
	defer func() { jwtKey = originalJwtKey }() // Restore original key

	t.Run("Successful Login", func(t *testing.T) {
		models.Users = make(map[string]models.User)
		models.UserEmailIndex = make(map[string]string)
		testEmail := "login@example.com"
		testPassword := "password123"
		testName := "Login User"
		_, err := setupTestUserForLogin(testEmail, testPassword, testName)
		assert.NoError(t, err, "Setup test user should not fail")

		payload := LoginRequest{Email: testEmail, Password: testPassword}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		LoginHandler(rr, req) // Call directly

		assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200 OK")

		var resp LoginResponse
		err = json.NewDecoder(rr.Body).Decode(&resp)
		assert.NoError(t, err, "Should decode response body without error")
		assert.NotEmpty(t, resp.Token, "Token should not be empty")
		assert.Equal(t, testEmail, resp.User.Email)
		assert.Equal(t, testName, resp.User.Name)

		// Optionally, verify token
		token, err := jwt.ParseWithClaims(resp.Token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		assert.NoError(t, err, "Token should be parseable")
		assert.True(t, token.Valid, "Token should be valid")
		claims, ok := token.Claims.(*Claims)
		assert.True(t, ok, "Token claims should be of type Claims")
		assert.Equal(t, testEmail, claims.Email)
	})

	t.Run("Invalid Password", func(t *testing.T) {
		models.Users = make(map[string]models.User)
		models.UserEmailIndex = make(map[string]string)
		testEmail := "wrongpass@example.com"
		_, err := setupTestUserForLogin(testEmail, "correctpassword", "Wrong Pass User")
		assert.NoError(t, err)

		payload := LoginRequest{Email: testEmail, Password: "incorrectpassword"}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		LoginHandler(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code, "Status code should be 401 Unauthorized")
	})

	t.Run("User Not Found", func(t *testing.T) {
		models.Users = make(map[string]models.User)
		models.UserEmailIndex = make(map[string]string)

		payload := LoginRequest{Email: "nonexistent@example.com", Password: "password123"}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		LoginHandler(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code, "Status code should be 401 Unauthorized")
	})

	t.Run("Missing Fields", func(t *testing.T) {
		payload := LoginRequest{Email: "test@example.com"} // Missing password
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		LoginHandler(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code, "Status code should be 400 Bad Request")
	})
}
