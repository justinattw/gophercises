package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type QA struct {
	Question string
	Answer	 string
	Response string
}

func (qa QA) String() string {
	return fmt.Sprintf("Q: %s, A: %s", qa.Question, qa.Answer)
}

func main() {

	// Defaults
	filename := "problems.csv"
	shuffle := false
	time := 30  // runtime of quiz in seconds
	time++
	
	// ----- Command line args -----
	if len(os.Args) >= 4 {
		newTime, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Printf("USAGE: go run quiz.go <filename> <shuffle quiz> <runtime in seconds>")
		} else {
			time = newTime
		}
	}
	if len(os.Args) >= 3 {
		if os.Args[2] == "t" {
			shuffle = true
		} else if os.Args[2] != "f" {
			log.Printf("USAGE: go run quiz.go <filename> <shuffle quiz: {t / f}> <runtime in seconds>")
		}
	}
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	}

	// ----- Set up questions list -----
	qaList, err := readCSV(filename)
	if err != nil {
		log.Printf("Failed to read file %s: %v\n", os.Args[1], err)
	}

	if shuffle {
		shuffleQuiz(qaList)
	}

	// ----- Run quiz -----
	responses := make(chan String)


	result := partOne(qaList)

	// Handle results
	printResult(result)
}

func partOne(qaList []QA) float64 {

	count := 0.0
	correct := 0.0

	for _, qa := range qaList {
		count++
		if ask(qa) {
			correct++
		}
	}

	result := (correct / count) * 100
	return result
}

func partTwo(qaList []QA) float64 {
	return 0.0
}

func ask(qa QA) bool {

	// Ask
	questionString := fmt.Sprintf("> ", qa.Question)
	fmt.Println(questionString)

	// Answer
	var input string
	fmt.Scan(&input)

	// Judge
	return input == qa.Answer
}



