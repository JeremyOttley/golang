package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  StdioHandling()
}

func StdioHandling() {
  s := bufio.NewScanner(os.Stdin)
  for s.Scan() {
    fmt.Println(s.Text())
  }
}
