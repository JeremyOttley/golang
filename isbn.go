package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"log"

	"golang.design/x/clipboard"
)

func main() {

	fmt.Println("Pease enter an isbn: ")
	reader := bufio.NewReader(os.Stdin)
	isbn, _ := reader.ReadString('\n')

	err := clipboard.Init()
	if err != nil {
        	log.Fatal(err)
	}

	r, _ := regexp.Compile("-")

	fmt.Println(r.ReplaceAllString(isbn, ""))
	clipboard.Write(clipboard.FmtText, []byte(r.ReplaceAllString(isbn, "")))

}
