package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Infinite loop that never ends
	for {
		// Prompt the user to enter a sentence
		fmt.Println("Enter a sentence:")

		// Create a reader to take input from the user
		reader := bufio.NewReader(os.Stdin)

		// Read the input line from the user until they press Enter
		text, _ := reader.ReadString('\n')

		// Remove extra spaces or newline characters from the input
		text = strings.TrimSpace(text)

		// Split the sentence into words
		words := strings.Fields(text)

		// Count the number of words
		wordCount := len(words)

		// Display the word count
		fmt.Printf("Word Count: %d\n", wordCount)
	}
}
