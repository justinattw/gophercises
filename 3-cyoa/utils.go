package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func readJson(filename string) (map[string]Gopher, error) {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %s\n", err)
		return nil, fmt.Errorf("Could not open file: %v", err)
	}

	defer file.Close()

	// Read the file's content
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	gopherMap := make(map[string]Gopher)
	err = json.Unmarshal(data, &gopherMap)
	if err != nil {
		log.Fatalf("Error unmarshaling json: %v", err)
		return nil, err
	}

	return gopherMap, nil
}
