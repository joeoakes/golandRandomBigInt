package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// Define the range for the random number.
	max := big.NewInt(100)

	// Generate a cryptographically secure random number.
	randomNumber, err := rand.Int(rand.Reader, max)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the random number.
	fmt.Println("Cryptographically secure random number for blockchain:", randomNumber)
}
