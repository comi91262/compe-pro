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
	var k = make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &k[i])
	}

	fmt.Fprintf(writer, "%v", k[0])
	for i := 0; i < n-2; i++ {
		fmt.Fprintf(writer, " ")
		fmt.Fprintf(writer, "%v", min(k[i], k[i+1]))
	}
	fmt.Fprintf(writer, " %v", k[n-2])
	fmt.Fprintf(writer, "\n")

}
