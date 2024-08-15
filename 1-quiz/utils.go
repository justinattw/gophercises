package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"math/rand"
	"strconv"
	"time"
)

// Read in CSV file into a slice of Quizzes
func readCSV(filename string) ([]Card, error) {
	
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %s\n", err)
		return nil, fmt.Errorf("Could not open file: %v", err)
	}
	
	defer file.Close()

	// Store file in data slice
	reader := csv.NewReader(file)

	// Construct question answers
	var cardsList []Card
	
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("Error while reading CSV record: %v", err)
		}
		if len(record) != 2 {
			return nil, fmt.Errorf("Unexpected number of fields for record: %v", record)
		}

		cardsList = append(cardsList, Card{record[0], record[1]})
	}

	return cardsList, nil
}

// Shuffle a QA slice
func shuffleQuiz[T any](slice []T) {
	// Seed the random number generator
    rand.Seed(time.Now().UnixNano())

    // Fisher-Yates shuffle algorithm
    n := len(slice)
    for i := n - 1; i > 0; i-- {
        j := rand.Intn(i + 1)
        slice[i], slice[j] = slice[j], slice[i]
    }
}

// Process results
func processResults(correct int, count int) float64 {
	var result float64
	if count == 0 {
		result = 0.0
	} else {
		result = (float64(correct) / float64(count)) * 100
	}
	printResult(result)
	return result
}

// Print results
func printResult(result float64) {
	fmt.Println()
	resultString := fmt.Sprintf("%s%% correct", strconv.FormatFloat(result, 'f', 2, 64))
	fmt.Println(resultString)
}

// Timer - sleeps for s seconds, then reports that it has finished
func timer(seconds int, done chan<- bool, cancel <-chan bool) {	
	select {
	case <- time.After(time.Duration(seconds) * time.Second):
		done <- true
		fmt.Println("\nTime's up!")
		return
	case <- cancel:
		return
	}
}

// Ask question
func ask(card Card, response chan<- bool) bool {

	questionString := fmt.Sprintf("> %s", card.Question)  // Ask
	input := waitForInput(questionString)  // Response
	res := input == card.Answer
	response <- res

	return res
}

func waitForInput(prompt string) string {
	fmt.Println(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}