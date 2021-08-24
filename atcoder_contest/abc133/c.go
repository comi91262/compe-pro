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

	var l int
	fmt.Fscan(reader, &l)
	var r int
	fmt.Fscan(reader, &r)

	ans := 1 << 60
	for i := l; i < r; i++ {
		for j := i + 1; j <= r; j++ {
			ans = min(ans, (i*j)%2019)
			if (i*j)%2019 == 0 {
				fmt.Fprintf(writer, "%v\n", 0)
				return
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
