package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")

	// Get all environment variables
	envVars := os.Environ()

	// Iterate and print each environment variable
	for _, env := range envVars {
		fmt.Println(env)
	}

	// Optionally, you can parse them into key-value pairs
	fmt.Println("\nParsed Key-Value Pairs:")
	for _, env := range envVars {
		pair := strings.SplitN(env, "=", 2) // Split only on the first '='
		if len(pair) == 2 {
			fmt.Printf("Key: %s, Value: %s\n", pair[0], pair[1])
		} else if len(pair) == 1 {
			// Handle cases where a variable might have no value (e.g., "KEY=")
			fmt.Printf("Key: %s, Value: (empty)\n", pair[0])
		}
	}
}
