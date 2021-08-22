package main

import (
	"fmt"
)

func main() {
	defer fmt.Print("ld.\n")

	defer fmt.Print("wor")

	fmt.Print("hello ")
}
