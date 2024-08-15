package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

type Card struct {
	Question string
	Answer	 string
}

func (card Card) String() string {
	return fmt.Sprintf("Q: %s, A: %s", card.Question, card.Answer)
}

func main() {

	// ----- Defaults and command line args -----
	filename := "problems.csv"
	shuffle := true
	timedMode := true
	seconds := 10
		
	if len(os.Args) >= 5 {
		if os.Args[4] == "t" {
			timedMode = true
		} else if os.Args[2] != "f" {
			log.Printf("USAGE: go run . <filename> <shuffle quiz: {t / f}> <runtime in seconds> <timedMode>")
			return
		}
	}
	if len(os.Args) >= 4 {
		newSeconds, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Printf("USAGE: go run . <filename> <shuffle quiz: {t / f}> <runtime in seconds> <timedMode>")
			return
		}
		seconds = newSeconds
	}
	if len(os.Args) >= 3 {
		if os.Args[2] == "t" {
			shuffle = true
		} else if os.Args[2] != "f" {
			log.Printf("USAGE: go run . <filename> <shuffle quiz: {t / f}> <runtime in seconds> <timedMode>")
			return
		}
	}
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	}

	// ----- Set up questions list -----
	cardsList, err := readCSV(filename)
	if err != nil {
		log.Printf("Failed to read file %s: %v\n", os.Args[1], err)
		return
	}

	if shuffle {
		shuffleQuiz(cardsList)
	}

	// ----- Run quiz -----
	if !timedMode {
		partOne(cardsList)
	} else {
		partTwo(cardsList, seconds)
	}
}

func partOne(cardsList []Card) {

	fmt.Println("You are playing untimed mode.")

	correct := 0
	total := len(cardsList)

	response := make(chan bool)
	defer close(response)
	
	for _, card := range cardsList {
		go ask(card, response)
		if <- response {
			correct++
		}
	}

	processResults(correct, total)
}

func partTwo(cardsList []Card, seconds int) {

	fmt.Printf("You are playing timed mode. You have %s seconds.\n", strconv.Itoa(seconds))
	
	correct := 0
	total := len(cardsList)
	
	responses := make(chan bool)
	timerDone := make(chan bool)
	timerCancel := make(chan bool)

	defer close(responses)
	defer close(timerDone)
	defer close(timerCancel)
	
	var wg sync.WaitGroup
	wg.Add(1)

	waitForInput("Press <Enter> to start.")

	go timer(seconds, timerDone, timerCancel)  // Start timer in a goroutine

	// Ask questions in a goroutine
	go func() {
		defer wg.Done()
		for _, card := range cardsList {

			go ask(card, responses)  // Ask question in a goroutine

			select {
			case <-timerDone:
				return
			case res := <-responses:
				if res {
					correct++
				}
			}
		}
		timerCancel <- true  // Send flag to timer goroutine to cancel
	}()

	wg.Wait()

	processResults(correct, total)
}

