package main

import (
	"math"
)

var a = [10]int{1, 14, 32, 51, 51, 51, 243, 419, 750, 910}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func isOK(index, key int) bool {
	return a[index] >= key
}

func binarySearch(key int) int {
	ng := -1
	ok := len(a)

	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		if isOK(mid, key) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func isOK(a *[]int, index, key int) bool {
	return (*a)[index] >= key
}

func binarySearch(a *[]int, key int) int {
	ng := -1
	ok := len(*a)

	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		if isOK(a, mid, key) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

func main() {
}
