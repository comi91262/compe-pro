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

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var t = make([]int, n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t[i])
		fmt.Fscan(reader, &a[i])
	}

	an := 0
	bn := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			an = t[i]
			bn = a[i]
			continue
		}

		if t[i] == a[i] {
			an, bn = max(an, bn), max(an, bn)
			continue
		}

		cnt := max((an+t[i]-1)/t[i], (bn+a[i]-1)/a[i])
		an = t[i] * cnt
		bn = a[i] * cnt

	}

	fmt.Fprintf(writer, "%v\n", an+bn)
}
