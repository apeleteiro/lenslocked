package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	switch os.Args[1] {
	case "hash":
		hash(os.Args[2])
	case "compare":
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Printf("Usage: %s [hash|compare] <password> [hash]\n", os.Args[0])
	}
}

func hash(password string) {
	fmt.Printf("Hashing password...\n")
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error hashing password: %v\n", err)
		return
	}
	fmt.Printf("Hashed password: %s\n", string(hashedBytes))
}

func compare(password, hash string) {
	fmt.Printf("Comparing password...\n")
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Printf("Passwords do not match: %v\n", err)
		return
	}
	fmt.Printf("Passwords match!\n")
}
