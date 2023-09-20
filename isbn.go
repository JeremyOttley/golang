import (
	"fmt"
	"regexp"
)

const isbn = "978-1-4780-2737-9"

func main() {

	r, _ := regexp.Compile("-")

	fmt.Println(r.ReplaceAllString(isbn, ""))

}
