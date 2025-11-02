package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

func main() {
	length := flag.Int("length", 12, "Length of the password")
	includeSpecial := flag.Bool("special", false, "Include special characters")
	flag.Parse()
	password := generatePassword(*length, *includeSpecial)
	fmt.Printf("Generated Password: %s\n", password)
}
func generatePassword(length int, special bool) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if special {
		letters += "!@#$%^&*()-_=+"
	}
	password := make([]byte, length)
	for i := range password {
		randIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		password[i] = letters[randIndex.Int64()]
	}
	return string(password)
}
