package main

import (
	"log"
	"os"
)

func main() {
	// Open the CSV file
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %s\n", err)
		return
	}
	defer file.Close()

	// Store file in data slice
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

}
