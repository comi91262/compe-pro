package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func primeFactor(n int) map[int]int {
	var m = map[int]int{}

	for i := 2; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		for n%i == 0 {
			m[i]++
			n /= i
		}
	}
	if n != 1 {
		m[n]++
	}
	return m
}
func floorLog(x int) int {
	if x == 1 {
		return 0
	}
	r := 1
	prod := 2
	for x > prod {
		prod *= 2
		r++
	}

	return r
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	fn := 0
	m := primeFactor(n)
	for _, v := range m {
		fn += v
	}

	fmt.Fprintf(writer, "%v\n", floorLog(fn))
}
