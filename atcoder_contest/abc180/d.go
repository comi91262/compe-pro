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
func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func main() {
	defer writer.Flush()

	var x int
	fmt.Fscan(reader, &x)
	var y int
	fmt.Fscan(reader, &y)
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	exp := 0

	for x*pow(a, exp) < y/a {
		exp++
	}

	ans := 0
	for i := 0; i <= exp; i++ {
		s := x * pow(a, i)
		cnt := (y - 1 - s) / b
		ans = max(ans, cnt+i)
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
