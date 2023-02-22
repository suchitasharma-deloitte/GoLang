package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Open the file for reading
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Initialize the word count
	wordCount := 0

	// Specify the word to search for
	searchWord := "The"

	// Read each line of the file
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into words
		words := strings.Fields(line)

		// Count the number of occurrences of the search word in the line
		for _, word := range words {
			if strings.ToLower(word) == strings.ToLower(searchWord) {
				wordCount++
			}
		}
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Print the word count
	fmt.Printf("The word \"%s\" appears %d times in the file.\n", searchWord, wordCount)
}
