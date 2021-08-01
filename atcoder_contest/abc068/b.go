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

func calc(n int) int {
	r := 0
	for n > 0 {
		if n%2 == 1 {
			return r
		}

		r++
		n /= 2
	}

	return r
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	mx := 0
	idx := 0
	for i := 1; i <= n; i++ {
		tmp := calc(i)
		if mx <= tmp {
			mx = tmp
			idx = i
		}

	}
	fmt.Fprintf(writer, "%v\n", idx)

}
