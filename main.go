package main

import (
	"fmt"
	"strconv")

func main() {

  fizzbuzz := []string{ "fizzbuzz%.0s", "%s", "%s", "fizz%.0s", "%s", "buzz%.0s", "fizz%.0s",
        "%s", "%s", "fizz%.0s", "buzz%.0s", "%s", "fizz%.0s", "%s", "%s" };

  for i := 1; i < 100; i++ {

    fmt.Printf( fizzbuzz[ i % 15 ] + "\r\n", strconv.Itoa(i) )
  }
}
