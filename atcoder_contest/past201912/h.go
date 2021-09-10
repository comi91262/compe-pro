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
	var n int
	fmt.Fscan(reader, &n)

	var c = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}

	ion, ien := -1, -1
	on := 1 << 60
	en := 1 << 60
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			if on >= c[i] {
				on = c[i]
				ion = i
			}
		} else {
			if en >= c[i] {
				en = c[i]
				ien = i
			}
		}
	}

	ans := 0
	ot := 0
	et := 0
	var q int
	fmt.Fscan(reader, &q)
	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(reader, &t)
		switch t {
		case 1:
			var x, a int
			fmt.Fscan(reader, &x, &a)
			x--

			switch x {
			case ion:
				if on >= a {
					ans += a
					on -= a
				}
			case ien:
				if en >= a {
					ans += a
					en -= a
				}
			default:
				if x%2 != 0 {
					if c[x]-et >= a {
						ans += a
						c[x] -= a
					}
				} else {
					if c[x]-ot >= a {
						ans += a
						c[x] -= a
					}
				}
			}
		case 2:
			var a int
			fmt.Fscan(reader, &a)

			if on >= a {
				ans += a * ((n + 1) / 2)
				ot += a
				on -= a
			}
		case 3:
			var a int
			fmt.Fscan(reader, &a)

			if min(on, en) >= a {
				en -= a
				on -= a
				et += a
				ot += a
				ans += a * n
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
