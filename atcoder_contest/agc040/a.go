package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
func add(arg ...int) (sum int) {
	for i := range arg {
		sum += arg[i]
	}
	return
}

func main() {
	defer writer.Flush()
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	n := len(s) + 1
	a := make([]int, n)

	for i := 1; i < n; i++ {
		if s[i-1] == "<" {
			a[i] = a[i-1] + 1
		}
	}

	for i := 0; i < n-1; i++ {
		if s[n-i-2] == ">" {
			a[n-i-2] = max(a[n-i-2], a[n-i-1]+1)
		}
	}

	fmt.Fprintf(writer, "%v\n", add(a...))
}
