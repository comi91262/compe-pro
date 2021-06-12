func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func fact(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

func combination(n int, r int) int {
	return fact(n) / fact(r) * fact(n-r)
}

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func min(x, y, z int) int {
	var m = int(math.Min(float64(x), float64(y)))
	return int(math.Min(float64(m), float64(z)))
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
		r *= a
		x--
	}
	return r
}
