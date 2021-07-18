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

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func f(a [][]int, x, n, p int) int {
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] == -1 {
				d[i][j] = x
			} else {
				d[i][j] = a[i][j]
			}

		}
	}

	// warshall_floyd
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				d[j][k] = min(d[j][k], d[j][i]+d[i][k])
			}
		}
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if d[i][j] <= p {
				cnt++
			}
		}
	}

	return cnt
}

func main() {
	defer writer.Flush()

	var n, p, k int
	fmt.Fscan(reader, &n, &p, &k)

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}

	ok := 1 << 60
	ng := 0 //n * (n - 1) / 2
	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2
		cnt := f(a, mid, n, p)
		if k >= cnt {
			ok = mid
		} else {
			ng = mid
		}
	}

	ok1 := 1 << 60
	ng1 := 0
	for abs(ok1-ng1) > 1 {
		mid := (ok1 + ng1) / 2
		cnt := f(a, mid, n, p)

		if k > cnt {
			ok1 = mid
		} else {
			ng1 = mid
		}
	}

	if ok1-ok > 1<<59 {
		fmt.Fprintf(writer, "%v\n", "Infinity")
	} else {
		fmt.Fprintf(writer, "%v\n", ok1-ok)
	}

}
