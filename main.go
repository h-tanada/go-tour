package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	morePrevious := -1
	previous := 1
	return func() int {
		result := morePrevious + previous
		morePrevious = previous
		previous = result
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
