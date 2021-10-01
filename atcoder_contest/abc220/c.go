package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func add(arg ...int) (sum int) {
	for i := range arg {
		sum += arg[i]
	}
	return
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	var x int
	fmt.Fscan(reader, &x)

	s := add(a...)

	k := x / s * n
	st := x / s * s
	for i := 0; i < n; i++ {
		if st+a[i] > x {
			k += i + 1
			break
		}
		st += a[i]
	}
	fmt.Fprintf(writer, "%v\n", k)
}
