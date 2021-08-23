package main

import (
	"bufio"
	"fmt"
	"math"
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
	var d int
	fmt.Fscan(reader, &d)

	x := make([][]int, n)
	for i := 0; i < n; i++ {
		x[i] = make([]int, d)
		for j := 0; j < d; j++ {
			fmt.Fscan(reader, &x[i][j])
		}
	}

	var a = make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i == j {
				continue
			}
			for k := 0; k < d; k++ {
				a[i][j] += (x[i][k] - x[j][k]) * (x[i][k] - x[j][k])
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i == j {
				continue
			}
			r := int(math.Sqrt(float64(a[i][j])))
			if r*r == a[i][j] {
				ans++
			}
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
