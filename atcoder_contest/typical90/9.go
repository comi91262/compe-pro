package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func getDegArgument(x, y float64) float64 {
	return math.Atan2(y, x) * 180.0 / math.Pi
}

func min(arg ...float64) float64 {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func max(arg ...float64) float64 {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func lowerBound(value, boader float64) bool {
	return boader <= value
}

func binarySearch(a []float64, boader float64, criteria func(value, boader float64) bool) int {
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

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var x = make([]float64, n)
	var y = make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
		fmt.Fscan(reader, &y[i])
	}

	ans := 0.0
	for i := 0; i < n; i++ {
		args := make([]float64, 0, n-1)

		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			ex, ey := x[j]-x[i], y[j]-y[i]
			args = append(args, getDegArgument(ex, ey))
		}

		sort.Slice(args, func(i, j int) bool { return args[i] <= args[j] })

		for _, arg := range args {
			pairArg := arg
			if pairArg >= 0.0 {
				pairArg -= 180.0
			} else {
				pairArg += 180.0
			}

			idx := binarySearch(args, pairArg, lowerBound)

			var angle float64
			if idx == len(args) {
				angle = min(math.Abs(args[idx-1]-arg), 360.0-math.Abs(args[idx-1]-arg))
			} else {
				angle = min(math.Abs(args[idx]-arg), 360.0-math.Abs(args[idx]-arg))
			}
			ans = max(ans, angle)
		}

	}

	fmt.Fprintf(writer, "%v\n", ans)
}
