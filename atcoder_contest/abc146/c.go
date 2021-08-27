package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func digit(x int) int {
	r := 0
	for x > 0 {
		r++
		x /= 10
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

func main() {
	defer writer.Flush()
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var x int
	fmt.Fscan(reader, &x)

	ok := -1
	ng := 1_000_000_001

	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		if a*mid+b*digit(mid) <= x {
			ok = mid
		} else {
			ng = mid
		}
	}
	fmt.Fprintf(writer, "%v\n", ok)

}
