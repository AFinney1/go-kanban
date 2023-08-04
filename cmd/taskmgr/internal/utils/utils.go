// Package for any general utilities (parsing dates, CLI helpers, etc).

/*
Parse command line flags - there are libraries like flag and pflag to help with this
Get user input - fmt.Scanln
Format output - fmt.Printf etc.
Dates - time.Parse to parse dates, time.Now to get current date
File I/O - ioutil to read/write files
Environment variables - os.Getenv
Logging - log package
Random numbers - math/rand
String manipulation - strings package
Exiting the app - os.Exit


*/

package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func GetUserInput(prompt string) string {
	// Use fmt.Printf and fmt.Scanln to get user input
	fmt.Printf(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetNow() time.Time {
	return time.Now()
}

func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}
