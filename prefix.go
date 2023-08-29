package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/antchfx/xmlquery"
)

func main() {

	fmt.Println("Enter a prefix: ")

	reader := bufio.NewReader(os.Stdin)

	prefix, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	requestUrl := fmt.Sprintf("https://doi.crossref.org/getPrefixPublisher/?prefix=%s", strings.TrimSpace(prefix))

	resp, err := http.Get(requestUrl)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	string_body := string(body)

	doc, err := xmlquery.Parse(strings.NewReader(string_))
	if err != nil {
		panic(err)
	}

	data := xmlquery.FindOne(doc, "//xml/publisher")

	if n := data.SelectElement("publisher_name"); n != nil {
		fmt.Println(n.InnerText())
	}
}
