package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func generateAPIKey(length int) (string, error) {
	key := make([]byte, length)

	_, err := rand.Read(key)

	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}
