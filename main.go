package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

var storedHash string

func registerUser(password string) {
	storedHash = generateHash(password)
	fmt.Println("Registration successful! SHA-256 hash code stored:", storedHash)
}

func generateHash(password string) string {
	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}

func authenticateUser(password string) bool {
	hash := generateHash(password)
	return storedHash == hash
}

func main() {
	fmt.Println("-Register a new user-")
	fmt.Println("Enter a password:")
	var password string
	fmt.Scanln(&password)
	registerUser(password)

	fmt.Println("\n-Authenticate the user-")
	for {
		fmt.Println("Enter your password (type '-exit' to exit):")
		fmt.Scanln(&password)

		if strings.TrimSpace(password) == "-exit" {
			break
		}

		if authenticateUser(password) {
			fmt.Println("Authentication successful!")
			break
		} else {
			fmt.Println("Authentication failed! Incorrect password.")
		}
	}
}
