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

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)

	var l = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &l[i])
	}

	power := make([]int, n+1)
	base := 4

	power[0] = 1
	for i := 0; i < n; i++ {
		power[i+1] = power[i] * base
	}

	ans := 1 << 60
	for i := 0; i < power[n]; i++ {
		aa := 0
		bb := 0
		cc := 0
		mp := 0
		for j := 0; j < n; j++ {
			bit := i / power[j] % base
			switch bit {
			case 0:
				if aa != 0 {
					mp += 10
				}
				aa += l[j]
			case 1:
				if bb != 0 {
					mp += 10
				}
				bb += l[j]
			case 2:
				if cc != 0 {
					mp += 10
				}
				cc += l[j]
			case 3:
			}
		}
		if aa == 0 || bb == 0 || cc == 0 {
			continue
		}
		mp += abs(a - aa)
		mp += abs(b - bb)
		mp += abs(c - cc)

		ans = min(mp, ans)
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
