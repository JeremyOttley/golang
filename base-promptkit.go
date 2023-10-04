package main

import (
	"fmt"
	"os"
	"log"
	
	"github.com/erikgeiser/promptkit/textinput"
)

func main() {

	input := textinput.New("Please enter text: ")
	input.Placeholder = "Text"
	
	userInput, err := input.RunPrompt()	
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		log.Fatal(err)
		os.Exit(1)
	}
	
	fmt.Println("Here's your text back: ", userInput)
	
}
