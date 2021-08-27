package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func toDigits(x, base int) []int {
	if x == 0 {
		return []int{0}
	}

	r := make([]int, 0)
	for x != 0 {
		r = append(r, x%base)
		x = x / base
	}

	return r
}
func contains(x int, a []int) bool {
	for _, e := range a {
		if x == e {
			return true
		}
	}
	return false
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)

	var d = make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &d[i])
	}

	for {
		digits := toDigits(n, 10)
		flag := false
		for i := range d {
			if contains(d[i], digits) {
				flag = true
			}
		}
		if !flag {
			fmt.Fprintf(writer, "%v\n", n)
			break
		}
		n++
	}

}
