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
	var k int
	fmt.Fscan(reader, &k)
	var m int
	fmt.Fscan(reader, &m)

	var a = make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &a[i])
	}

	sum := 0
	for i := range a {
		sum += a[i]
	}

	ans := max(m*n-sum, 0)

	if ans <= k {
		fmt.Fprintf(writer, "%v\n", ans)
	} else {
		fmt.Fprintf(writer, "%v\n", -1)
	}

}
