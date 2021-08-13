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
	var h = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &h[i])
	}

	mh := max(h...)
	ans := 0
	for i := 0; i < mh; i++ {
		// fmt.Fprintf(writer, "%v\n", h)
		m := 0
		for j := 0; j < n; j++ {
			if h[j] != 0 {
				continue
			}
			if j-1 >= 0 {
				if h[j-1] != 0 {
					m++
				}
			}

			if j+1 < n {
				if h[j+1] != 0 {
					m++
				}
			}
		}

		if h[0] != 0 {
			m++
		}
		if h[n-1] != 0 {
			m++
		}
		// fmt.Fprintf(writer, "%v\n", m/2)
		ans += m / 2

		// fmt.Fprintf(writer, "%v\n", ans)
		for j := 0; j < n; j++ {
			if h[j] != 0 {
				h[j]--
			}
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
