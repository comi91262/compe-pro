package algo

import "math"

func getDegArgument(x, y float64) float64 {
	return math.Atan2(y, x) * 180.0 / math.Pi
}

func getRadArgument(x, y float64) float64 {
	return math.Atan2(y, x)
}

// 時計回り
func rotate(x, y, deg float64) (float64, float64) {
	cs := math.Cos(deg * math.Pi / 180.0)
	sn := math.Sin(deg * math.Pi / 180.0)

	return x*cs - y*sn, x*sn + y*cs
}
