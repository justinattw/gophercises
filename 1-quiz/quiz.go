package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type QA struct {
	Number   int
	Question string
	Answer	 string
}

func (qa QA) String() string {
	return fmt.Sprintf("#%s, Q: %s, A: %s", strconv.Itoa(qa.Number), qa.Question, qa.Answer)
}

func main() {
	data := readCSV()
	qaList := constructQuestionAnswers(data)
	fmt.Println()

	// Run quiz
	result := partOne(qaList)

	// Print result
	printResult(result)
}

func partOne(qaList []QA) float64 {

	count := 0.0
	correct := 0.0

	for i, qa := range qaList {
		questionString := fmt.Sprintf("Question #%s: %s = ", strconv.Itoa(i+1), qa.Question)
		fmt.Print(questionString)

		var input string
		fmt.Scan(&input)

		if input == qa.Answer {
			correct++
		}
		count++
	}

	result := (correct / count) * 100

	return result
}

func readCSV() [][]string {
	
	// Open the CSV file
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %s\n", err)
	}
	
	defer file.Close()

	// Store file in data slice
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func constructQuestionAnswers(data [][]string) []QA {
	var qaList []QA
	
	for i, line := range data {
		var rec QA
		rec.Number = i + 1
		rec.Question = line[0]
		rec.Answer = line[1]
		
		fmt.Println(rec)
		qaList = append(qaList, rec)
	}

	return qaList
}

func printResult(result float64) {
	fmt.Println()
	resultString := fmt.Sprintf("%s%% correct", strconv.FormatFloat(result, 'f', 2, 64))
	fmt.Println(resultString)
}