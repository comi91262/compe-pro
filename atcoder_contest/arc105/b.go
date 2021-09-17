package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func maxElement(arg ...int) int {
	max := arg[0]
	me := 0
	for i, x := range arg {
		if max < x {
			max = x
			me = i
		}
	}
	return me
}
func minElement(arg ...int) int {
	min := arg[0]
	me := 0
	for i, x := range arg {
		if min > x {
			min = x
			me = i
		}
	}
	return me
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	i := maxElement(a...)
	j := minElement(a...)
	for a[i] != a[j] {
		for k := 0; k < n; k++ {
			a[k] = gcd(a[k], a[j])
		}
		i = maxElement(a...)
		j = minElement(a...)
	}

	fmt.Fprintf(writer, "%v\n", a[i])
}
