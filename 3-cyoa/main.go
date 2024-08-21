package main

import "fmt"

type Gopher struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func main() {

	// 1. Use the html/template package to create your HTML pages. Part of the purpose of this exercise is to get practice using this package.

	// 2. Create an http.Handler to handle the web requests instead of a handler function.

	// 3. Use the encoding/json package to decode the JSON file. You are welcome to try out third party packages afterwards, but I recommend starting here.
	storyMap, err := readJson("gopher.json")
	if err != nil {
		return
	}

	// Print the map
	for location, gopher := range storyMap {
		fmt.Printf("Key: %s, Title: %s\n", location, gopher.Story)
	}

}
