package auth

import (
	"encoding/json"
	"net/http"
	"time"
	"wallet-app/backend/models" // Ensure this path is correct based on your go.mod
	"github.com/golang-jwt/jwt/v5"
)

// Define a secret key for signing JWT tokens.
// WARNING: In a production application, this should be loaded from a secure configuration or environment variable.
var jwtKey = []byte("my_super_secret_key_that_should_be_very_long_and_random")

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  struct {
		UserID string `json:"userId"`
		Email  string `json:"email"`
		Name   string `json:"name"`
	} `json:"user"`
}

// Claims struct for JWT
type Claims struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Find user by email
	userID, emailExists := models.UserEmailIndex[req.Email]
	if !emailExists {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized) // Generic message for security
		return
	}

	user, userExists := models.Users[userID]
	if !userExists {
		// This case should ideally not happen if UserEmailIndex is consistent with Users map
		http.Error(w, "User data integrity issue", http.StatusInternalServerError)
		return
	}

	// Check password
	if !models.CheckPasswordHash(req.Password, user.Password) {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized) // Generic message
		return
	}

	// Generate JWT token
	expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours
	claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "wallet-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	resp := LoginResponse{
		Token: tokenString,
		User: struct {
			UserID string `json:"userId"`
			Email  string `json:"email"`
			Name   string `json:"name"`
		}{
			UserID: user.ID,
			Email:  user.Email,
			Name:   user.Name,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
