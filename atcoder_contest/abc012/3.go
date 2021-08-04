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

	var n int
	fmt.Fscan(reader, &n)

	sum := 0
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			sum += i * j
		}
	}

	d := divisor(sum - n)

	sort.Ints(d)
	for i := range d {
		if 9 < d[i] {
			continue
		}
		if 9 < (sum-n)/d[i] {
			continue
		}
		fmt.Fprintf(writer, "%v x %v\n", d[i], (sum-n)/d[i])
	}
}
