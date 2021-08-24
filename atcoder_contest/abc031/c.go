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
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := -1 << 60
	for i := 0; i < n; i++ {
		omax := -1 << 60
		tmax := -1 << 60
		for j := 0; j < n; j++ {
			o := 0
			t := 0
			if i == j {
				continue
			}
			l := min(i, j)
			r := max(i, j)

			cnt := 0
			for k := l; k <= r; k++ {
				if cnt%2 != 0 {
					o += a[k]
				} else {
					t += a[k]
				}
				cnt++
			}
			if omax < o {
				omax = o
				tmax = t
			}
		}

		ans = max(ans, tmax)
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
