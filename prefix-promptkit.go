package main

import (
	"fmt"
	"log"
	"os"

	"github.com/antchfx/xmlquery"
	"github.com/erikgeiser/promptkit/textinput"
)

func main() {
	input := textinput.New("Please enter a prefix: ")
	input.Placeholder = "10.1215"

	prefix, err := input.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	crossrefApi := fmt.Sprintf("https://doi.crossref.org/getPrefixPublisher/?prefix=%s", prefix)

	doc, err := xmlquery.LoadURL(crossrefApi)
	if err != nil {
		log.Fatalln(err)
	}

	data := xmlquery.FindOne(doc, "//xml/publisher")

	if n := data.SelectElement("publisher_name"); n != nil {
		fmt.Println(n.InnerText())
	}
}
