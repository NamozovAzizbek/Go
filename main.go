package main

import (
	"fmt"

	"strconv"
)
func main() {
	var n int = 123
    t := strconv.Itoa(n)
    fmt.Printf("%t", t)
}