func pow(a, x int) int {
	r := 1
	for x > 0 {
		r *= a
		x--
	}
	return r
}
