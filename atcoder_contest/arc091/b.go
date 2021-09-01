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

	ans := 0
	for i := 1; i <= n; i++ {
		ans += n / i * max(i-k, 0)
		ans += max(0, n%i-k+1)
	}
	cnt := 0
	for i := 1; i <= n; i++ {
		if 0/i >= k {
			cnt++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans-cnt)
}
