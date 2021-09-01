package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

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

func main() {
	defer writer.Flush()
	var q int
	fmt.Fscan(reader, &q)
	var l = make([]int, q)
	var r = make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &l[i])
		fmt.Fscan(reader, &r[i])
	}

	size := 100001
	p := makePrimes(size)
	b := make([]int, size)
	for i := 0; i < len(p); i++ {
		b[p[i]]++
	}
	c := make([]int, size)
	copy(c, b)
	for i := 1; i < size; i++ {
		if b[i] != 1 {
			continue
		}
		if b[(i+1)/2] != 1 {
			c[i] = 0
		}
	}

	for i := 0; i < size-1; i++ {
		c[i+1] += c[i]
	}

	for i := 0; i < q; i++ {
		fmt.Fprintf(writer, "%v\n", c[r[i]]-c[l[i]-1])
	}

}
