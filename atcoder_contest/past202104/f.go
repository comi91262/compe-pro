package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var l int
	fmt.Fscan(reader, &l)
	var t int
	fmt.Fscan(reader, &t)
	var x int
	fmt.Fscan(reader, &x)
	var a = make([]int, n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
	}

	for i := 0; i < n; i++ {
		if b[i] >= l && a[i] > t {
			fmt.Fprintf(writer, "%v\n", "forever")
			return
		}

	}

	ans := 0
	tt := 0
	for i := 0; i < n; i++ {
		if b[i] < l {
			tt = 0
			ans += a[i]
			continue
		}

		tt += a[i]
		for tt >= t {
			ans += x
			if tt == t {
				tt = 0
				break
			}
			ans += t - tt + a[i]
			tt = a[i]
			if b[i] >= l && a[i] > t {
				fmt.Fprintf(writer, "%v\n", "forever")
				return
			}
		}

		ans += a[i]
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
