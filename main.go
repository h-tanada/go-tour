package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	splitS := strings.Split(s, " ")
	result := make(map[string]int)
	for _, value := range splitS {
		result[value]++
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
