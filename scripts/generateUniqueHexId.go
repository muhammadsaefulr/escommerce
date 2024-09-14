package scripts

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateUniqueHexId6B() (string, error) {
	bytes := make([]byte, 6)
	_, err := rand.Read(bytes)

	return hex.EncodeToString(bytes), err
}

func GenerateUniqueHexId4B() (string, error) {
	bytes := make([]byte, 4)
	_, err := rand.Read(bytes)

	return hex.EncodeToString(bytes), err
}

func GenerateUniqueHexId3B() (string, error) {
	bytes := make([]byte, 3)
	_, err := rand.Read(bytes)

	return hex.EncodeToString(bytes), err
}
