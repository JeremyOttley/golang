package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"golang.design/x/clipboard"
)

func main() {

	fmt.Println("Pease enter an isbn: ")
	reader := bufio.NewReader(os.Stdin)
	isbn, _ := reader.ReadString('\n')

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	r, _ := regexp.Compile("-")

	fmt.Println(r.ReplaceAllString(isbn, ""))
	clipboard.Write(clipboard.FmtText, []byte(r.ReplaceAllString(isbn, "")))

}
