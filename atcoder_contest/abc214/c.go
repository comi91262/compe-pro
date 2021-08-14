package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var s = make([]int, n)
	var t = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t[i])
	}

	mn := 1 << 60
	start := 0
	for i := range t {
		if mn > t[i] {
			mn = t[i]
			start = i
		}
	}
	// fmt.Fprintf(writer, "%v\n", start)

	var ss = make([]int, 2*n)
	var tt = make([]int, 2*n)
	for i := 0; i < n; i++ {
		ss[i], ss[i+n] = s[i], s[i]
		tt[i], tt[i+n] = t[i], t[i]
	}
	// fmt.Fprintf(writer, "%v\n", ss)
	// fmt.Fprintf(writer, "%v\n", tt)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		a[(start+i+n)%n] = min(a[(start+i-1+n)%n]+ss[(start+i-1+n)%n], tt[start+i])
	}

	// fmt.Fprintf(writer, "%v\n", a)

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%v\n", a[i])
	}

}
