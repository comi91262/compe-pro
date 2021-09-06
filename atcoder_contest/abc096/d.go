package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func makePrimes(n int) []int {
	r := []int{}
	prime := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		prime[i] = 1
	}

	prime[0], prime[1] = 0, 0
	for p := 2; p*p < n; p++ {
		if prime[p] > 0 {
			for x := 0; p*(x+2) <= n; x++ {
				prime[p*(x+2)] = 0
			}
		}
	}
	for p := 2; p < n+1; p++ {
		if prime[p] > 0 {
			r = append(r, p)
		}
	}

	return r
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}
func lowerBound(value, boader int) bool {
	return boader <= value
}

func binarySearch(a []int, boader int, criteria func(value, boader int) bool) int {
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

	rand.Seed(time.Now().UnixNano())

	var n int
	fmt.Fscan(reader, &n)

	p := makePrimes(55555)
	ans := []int{}
	for i := 0; i < len(p); i++ {
		if p[i]%5 == 1 {
			ans = append(ans, p[i])
		}

		if len(ans) == n {
			break
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%v", ans[i])
		if i != n-1 {
			fmt.Fprintf(writer, " ")
		}

	}
	fmt.Fprintf(writer, "\n")
}
