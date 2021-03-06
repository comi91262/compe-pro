package algo

// 答えで二分探索用
func binarySearch() int {
	l := -1
	r := 1 << 60

	for r-l > 1 {
		mid := (l + r) / 2

		f := func(x int) bool {
			return true
		}

		if f(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

func bit2FullSearch(n int) {
	for i := 0; i < 1<<n; i++ {
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
			}
		}
	}
}

func bitfullSearch(base int, n int) {
	power := make([]int, n+1)
	power[0] = 1
	for i := 0; i < n; i++ {
		power[i+1] = power[i] * base
	}

	for i := 0; i < power[n]; i++ {
		for j := 0; j < n; j++ {
			bit := i / power[j] % base
			switch bit {
			case 0:
				// ...
			}
		}
	}
}

func ternarySearch(n int, a []int) int {
	ans, cl, cr := 0, 1, n

	for i := 0; i < n; i++ {
		c1 := (cl + cl + cr) / 3
		c2 := (cl + cr + cr) / 3

		d1 := a[c1]
		d2 := a[c2]

		ans = max(ans, d1, d2)
		if d1 < d2 {
			cl = c1
		} else {
			cr = c2
		}
	}
	return ans
}
