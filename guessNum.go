// Package for trying out the Go language
package goplayground

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// stripSpaces strips spaces (space, newlines, tabs) from a given string
// Implemented for writing funcs with param and return value
func stripSpaces(s string) string {
	return strings.TrimSpace(s)
}

// Guess guesses Number between 1 and 100 for 10 times
// Example from Head First Go
func Guess() {
	success := false
	epoch := time.Now().Unix()
	rand.Seed(epoch)
	target := rand.Intn(100) + 1 // [0, 99]
	for numGuesses := 0; numGuesses < 10; numGuesses++ {
		fmt.Print("Guess Num:\n>")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = stripSpaces(input)
		inputInt, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if target < inputInt {
			fmt.Println("Guess too big")
		} else if target > inputInt {
			fmt.Println("Guess too small")
		} else {
			fmt.Printf("You made it!\n")
			success = true
			break
		}
	}
	if !success {
		fmt.Printf("Out of guesses!\n")
	}
}
