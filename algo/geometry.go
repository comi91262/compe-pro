package algo

import "math"

func getDegArgument(x, y float64) float64 {
	return math.Atan2(y, x) * 180.0 / math.Pi
}

func getRadArgument(x, y float64) float64 {
	return math.Atan2(y, x)
}
