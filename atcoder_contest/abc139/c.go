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
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := 0
	pre := a[0]
	cnt := 1
	for i := 1; i < n; i++ {
		if pre < a[i] {
			ans = max(ans, cnt-1)
			cnt = 0
		}
		cnt++
		pre = a[i]
	}
	fmt.Fprintf(writer, "%v\n", max(ans, cnt-1))
}
