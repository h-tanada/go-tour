package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[:2]
	b := names[1:]
	fmt.Println(a, b)

	b[0] = "Corouse"
	fmt.Println(a, b)
	fmt.Println(names)
}
