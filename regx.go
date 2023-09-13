package main

import (
	"fmt"
	"regexp"
)

const PhoneNumber = "410-581-3653"

func main() {

	r, _ := regexp.Compile("(?P<areaCode>\\d{3}).\\d{3}.\\d{4}")
	fmt.Println(r.MatchString(PhoneNumber))
	fmt.Println(r.FindString(PhoneNumber))

	//match, _ := regexp.MatchString("(?P<areaCode>\\d{3}).\\d{3}.\\d{4}", PhoneNumber)
	//fmt.Println(match)
}
