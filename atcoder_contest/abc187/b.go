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

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var x = make([]int, n)
	var y = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
		fmt.Fscan(reader, &y[i])
	}

	ans := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			dy := abs(y[j] - y[i])
			dx := abs(x[j] - x[i])

			if dx >= dy {
				ans++
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
