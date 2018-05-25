package main

import (
	"fmt"
	"strconv"
)

func main() {
	x, err := strconv.Atoi("42")
	if err != nil {
		return
	}
	fmt.Println(x)
}
