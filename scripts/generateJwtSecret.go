package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateSecretKey() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(key)
}

func main() {
	secretKey := generateSecretKey()
	fmt.Println("Generated Secret Key:", secretKey)
}
