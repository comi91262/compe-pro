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
	var d = make([]int, 7)
	for i := 0; i < 7; i++ {
		fmt.Fscan(reader, &d[i])
	}
	var j = make([]int, 7)
	for i := 0; i < 7; i++ {
		fmt.Fscan(reader, &j[i])
	}

	ans := 0
	for i := 0; i < 7; i++ {
		ans += max(d[i], j[i])
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
