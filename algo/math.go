package algo

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	d := gcd(a, b)
	return a / d * b
}

func permutation(n int, k int) int {
	if n < k {
		return 0
	}

	prod := 1
	for k > 0 {
		prod *= n - k + 1
		k--
	}
	return prod
}

func combination(n int, k int) int {
	return permutation(n, k) / permutation(k, k)
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

// 約数列挙
func divisor(n int) []int {
	var r []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			r = append(r, i)
			if i*i != n {
				r = append(r, n/i)
			}
		}
	}
	return r
}
