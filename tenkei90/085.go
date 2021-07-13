package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

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

func main() {
	defer writer.Flush()

	var k int
	fmt.Fscan(reader, &k)

	d := divisor(k)

	sort.Ints(d)

	ans := 0
	for i := 0; i < len(d); i++ {
		for j := i; j < len(d); j++ {

			c := k / d[i] / d[j]
			if c*d[i]*d[j] != k {
				continue
			}

			if d[i] <= d[j] && d[j] <= c {
				ans++
			}

		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
