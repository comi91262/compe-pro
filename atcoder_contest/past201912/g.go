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

	a := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		a[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}

	power := make([]int, n+1)
	base := 3
	power[0] = 1

	for i := 0; i < n; i++ {
		power[i+1] = power[i] * base
	}

	ans := -1 << 60
	for i := 0; i < power[n]; i++ {
		b := make([][]int, base)

		for j := 0; j < n; j++ {
			bit := i / power[j] % base
			switch bit {
			case 0:
				b[0] = append(b[0], j)
			case 1:
				b[1] = append(b[1], j)
			case 2:
				b[2] = append(b[2], j)
			}
		}

		fav := 0
		for j := 0; j < base; j++ {
			m := len(b[j])
			if m <= 1 {
				continue
			}
			for k := 0; k < m; k++ {
				for l := k + 1; l < m; l++ {
					s := b[j][k]
					t := b[j][l]
					fav += a[s][t]
				}
			}
		}
		ans = max(ans, fav)

	}

	fmt.Fprintf(writer, "%v\n", ans)

}
