package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for y, _ := range result {
		result[y] = make([]uint8, dx)
		for x, _ := range result[y] {
			result[y][x] = calc(x, y)
		}
	}
	return result
}

func calc(x, y int) uint8 {
	return uint8(y ^ x)
}

func main() {
	pic.Show(Pic)
}
