package main

import (
        "fmt"
		"strconv"
)

var (
	cube = func(i int) string {
			c := i * i * i
			return strconv.Itoa(c)
	}
)

func main() {
	x := cube(8)
	fmt.Printf("%T %v", x, x)
}