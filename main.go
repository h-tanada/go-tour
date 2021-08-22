package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("when's Saturday?")
	today := time.Now().Weekday()
	switch time.Monday - today {
	case 0:
		fmt.Println("Today.")
	case 1:
		fmt.Println("Tomorrow.")
	case 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}
