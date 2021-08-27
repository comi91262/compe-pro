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

	a := make([]int, n)
	x := make([][]int, n)
	y := make([][]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		x[i] = make([]int, a[i])
		y[i] = make([]int, a[i])

		for j := 0; j < a[i]; j++ {
			fmt.Fscan(reader, &x[i][j])
			fmt.Fscan(reader, &y[i][j])
			x[i][j]--
		}
	}

	ans := 0
	for i := 0; i < 1<<n; i++ {

		isValid := true
		for j := 0; j < n; j++ {
			if i&(1<<j) == 0 {
				continue
			}
			for k := 0; k < a[j]; k++ {
				if y[j][k] == 1 && i&(1<<x[j][k]) == 0 {
					isValid = false
					break
				}

				if y[j][k] == 0 && i&(1<<x[j][k]) != 0 {
					isValid = false
					break
				}
			}

		}
		b := i
		if isValid {
			cnt := 0
			for b > 0 {
				cnt += b % 2
				b /= 2
			}
			ans = max(cnt, ans)
		}

	}
	fmt.Fprintf(writer, "%v\n", ans)
}
