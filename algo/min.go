package main

import "math"

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func min(x, y, z int) int {
	var m = int(math.Min(float64(x), float64(y)))
	return int(math.Min(float64(m), float64(z)))
}
