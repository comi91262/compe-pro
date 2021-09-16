package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}
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
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)

	ans := 1 << 60
	ans = min(ans, abs(a*b*(c/2)-a*b*((c+1)/2)))
	ans = min(ans, abs(a*c*(b/2)-a*c*((b+1)/2)))
	ans = min(ans, abs(c*b*(a/2)-c*b*((a+1)/2)))
	fmt.Fprintf(writer, "%v\n", ans)

}
