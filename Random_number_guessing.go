package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	target := rand.Intn(100) + 1     // Pick a number between 1 and 100
	var guess int

	fmt.Println("Guess the number between 1 and 100!")

	for {
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scanln(&guess)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if guess < target {
			fmt.Println("Too low. Try again.")
		} else if guess > target {
			fmt.Println("Too high. Try again.")
		} else {
			fmt.Println("🎉 Correct! You guessed it.")
			break
		}
	}
}
