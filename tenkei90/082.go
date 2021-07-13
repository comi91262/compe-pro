package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const d = 1_000_000_000 + 7

func digits(x uint) uint {
	var cnt uint
	for x > 0 {
		x /= 10
		cnt++
	}

	return cnt
}

func max(arg ...uint) uint {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func min(arg ...uint) uint {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func pow(a, x uint) uint {
	var r uint = 1
	for x > 0 {
		if x&1 == 1 {
			r *= a
		}
		a *= a
		x >>= 1
	}
	return r
}

func powMod(a, x, d uint) uint {
	var r uint = 1
	for x > 0 {
		if x&1 == 1 {
			r = r * a % d
		}
		a = a * a % d
		x >>= 1
	}
	return r
}

func calc(e, s uint) uint {
	t1 := (e - s) % d
	t2 := (e + s + 1) % d
	return t1 * t2 % d * powMod(2, d-2, d) % d
}

func main() {
	defer writer.Flush()

	var l, r uint
	fmt.Fscan(reader, &l, &r)

	nl := digits(l)
	nr := digits(r)

	var sum uint = 0
	if nl == nr {
		fmt.Fprintf(writer, "%v\n", nr*calc(r, l-1)%d)
		return
	}

	for i := nl; i <= nr; i++ {
		p := pow(10, i)
		q := pow(10, i-1)

		s := max(q-1, l-1)
		e := min(p-1, r)

		sum += i * calc(e, s) % d
	}

	fmt.Fprintf(writer, "%v\n", sum%d)
}
