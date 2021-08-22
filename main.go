package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	count := 0
	for math.Abs(math.Pow(z, 2)-x) > 0.0000000000001 {
		count++
		z -= calcurateNumberForAjustZ(x, z)
	}
	fmt.Println(count)
	return z
}
func calcurateNumberForAjustZ(x float64, z float64) float64 {
	return (math.Pow(z, 2) - x) / (z * 2)
}

func main() {
	for target := 1.0; target < 100; target++ {
		fmt.Println(Sqrt(target))
	}
}
