package main

import (
	"io/ioutil"
	"log"
	"net/http"
  "fmt"
  "encoding/json"
)

func main() {
  
	resp, err := http.Get("https://character.totalpartykill.ca/basic/json")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

  var reading map[string]interface{}
  err = json.Unmarshal([]byte(body), &reading)

  fmt.Printf("%+v\n", reading)

}
