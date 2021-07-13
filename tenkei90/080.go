package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

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

func main() {
	defer writer.Flush()

	var n, d int
	fmt.Fscan(reader, &n, &d)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := pow(2, d)
	for i := 0; i < 1<<n; i++ {
		c := []int{}
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				c = append(c, j)
			}
		}
		if len(c) == 0 {
			continue
		}

		bit := 0
		for _, idx := range c {
			bit |= a[idx]
		}

		cnt := 0
		for j := 0; j < d; j++ {
			if ((bit >> j) & 1) == 0 {
				cnt++
			}
		}

		if len(c)%2 == 0 {
			ans += pow(2, cnt)
		} else {
			ans -= pow(2, cnt)
		}

	}

	fmt.Fprintf(writer, "%v\n", ans)
}
