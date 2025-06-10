package models

import (
	"time"
	"golang.org/x/crypto/bcrypt"
)

// User defines the structure for a user
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Store hashed password, '-' means don't marshal to JSON
	CreatedAt time.Time `json:"createdAt"`
}

// HashPassword hashes the user's password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a plain password with a hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Temporary in-memory store for users
var Users = make(map[string]User) // Using a map for easier lookup by ID or email
var UserEmailIndex = make(map[string]string) // Email to UserID index
