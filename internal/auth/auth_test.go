package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	// Setup test data
	userID := uuid.New()
	tokenSecret := "test-secret"
	expiresIn := time.Hour // 1 hour from now

	// Create the JWT
	tokenString, err := MakeJWT(userID, tokenSecret, expiresIn)
	if err != nil {
		t.Fatalf("Failed to create JWT: %v", err)
	}

	// Validate the JWT
	parsedUserID, err := ValidateJWT(tokenString, tokenSecret)
	if err != nil {
		t.Fatalf("Failed to validate JWT: %v", err)
	}

	// Check that we got back the same user ID
	if parsedUserID != userID {
		t.Errorf("Expected user ID %v, got %v", userID, parsedUserID)
	}
}

func TestValidateJWT_Expired(t *testing.T) {
	userID := uuid.New()
	tokenSecret := "test-secret"
	expiresIn := -time.Hour // Already expired (1 hour ago)

	// Create an expired JWT
	tokenString, err := MakeJWT(userID, tokenSecret, expiresIn)
	if err != nil {
		t.Fatalf("Failed to create JWT: %v", err)
	}

	// Try to validate the expired JWT
	_, err = ValidateJWT(tokenString, tokenSecret)
	if err == nil {
		t.Error("Expected validation to fail for expired token, but it succeeded")
	}
}

func TestValidateJWT_WrongSecret(t *testing.T) {
	userID := uuid.New()
	correctSecret := "correct-secret"
	wrongSecret := "wrong-secret"
	expiresIn := time.Hour

	// Create JWT with correct secret
	tokenString, err := MakeJWT(userID, correctSecret, expiresIn)
	if err != nil {
		t.Fatalf("Failed to create JWT: %v", err)
	}

	// Try to validate with wrong secret
	_, err = ValidateJWT(tokenString, wrongSecret)
	if err == nil {
		t.Error("Expected validation to fail with wrong secret, but it succeeded")
	}
}
