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

// Print results
func printResult(result float64) {
	fmt.Println()
	resultString := fmt.Sprintf("%s%% correct", strconv.FormatFloat(result, 'f', 2, 64))
	fmt.Println(resultString)
}