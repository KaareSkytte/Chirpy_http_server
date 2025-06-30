package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func MakeRefreshToken() (string, error) {
	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	encodedToken := hex.EncodeToString(bytes)

	return encodedToken, nil
}
