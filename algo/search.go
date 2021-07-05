package algo

func upperBound(value, boader int) bool {
	return boader < value
}

func lowerBound(value, boader int) bool {
	return boader <= value
}

func binarySearch(a []int, boader int, criteria func(value, boader int) bool) int {
	abs := func(x int) int {
		if x >= 0 {
			return x
		} else {
			return x * -1
		}
	}

	ng := -1
	ok := len(a)

	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		if criteria(a[mid], boader) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

func bitFullSearch(n int) {
	for i := 0; i < 1<<n; i++ {
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
			}
		}
	}
}
