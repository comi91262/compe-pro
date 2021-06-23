package algo

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

// x = y のとき 一番左のものを返す (std::min)
func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func mod(x, y int) int {
	x = x % y
	if x >= 0 {
		return x
	}
	if y < 0 {
		return x - y
	}
	return x + y
}

func pow(a, x int) int {
	r := 1
	for x > 0 {
		if x&1 == 1 {
			r *= a
		}
		a *= a
		x >>= 1
	}
	return r
}
