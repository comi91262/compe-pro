package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func check(a []int, start bool) int {
	n := len(a)
	ans := 0
	sum := 0

	if start {
		if a[0] <= 0 {
			ans += abs(1 - a[0])
			a[0] = 1
		}
	} else {
		if a[0] >= 0 {
			ans += abs(-1 - a[0])
			a[0] = -1
		}
	}

	for i := 0; i < n-1; i++ {
		sum += a[i]
		if sum*(sum+a[i+1]) < 0 && sum != 0 {
			continue
		}
		if sum == 0 {
			sum++
			ans++
		}

		if sum > 0 {
			if sum+a[i+1] > 0 {
				ans += abs(-1 - (sum + a[i+1]))
				a[i+1] = -1 - sum
			}
		} else {
			if sum+a[i+1] < 0 {
				ans += abs(1 - (sum + a[i+1]))
				a[i+1] = 1 - sum
			}
		}
	}
	if sum+a[n-1] == 0 {
		ans++
	}
	return ans
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}
func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	b := make([]int, len(a))
	c := make([]int, len(a))
	copy(b, a)
	copy(c, a)

	fmt.Fprintf(writer, "%v\n", min(check(b, true), check(c, false)))
}
