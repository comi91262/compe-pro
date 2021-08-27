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

	var f = make([][10]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < 10; j++ {
			fmt.Fscan(reader, &f[i][j])
		}
	}

	var p = make([][11]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < 11; j++ {
			fmt.Fscan(reader, &p[i][j])
		}
	}

	ans := -1 << 60
	for i := 1; i < 1<<10; i++ {
		pr := 0
		for k := 0; k < n; k++ {
			cnt := 0
			for j := 0; j < 10; j++ {
				if i&(1<<j) != 0 && f[k][j] == 1 {
					cnt++
				}
			}
			pr += p[k][cnt]
		}

		ans = max(pr, ans)
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
