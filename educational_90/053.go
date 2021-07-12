package main

import (
	"fmt"
)

var fib = [16]int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597}
var n int
var idx [1609]int

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
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		solve()
	}

}

func ask(pos int) int {
	if idx[pos] == -1 {
		fmt.Printf("? %v\n", pos)
		var t int
		fmt.Scan(&t)
		idx[pos] = t
		return t
	}
	return idx[pos]
}

func solve() {
	var n int
	fmt.Scan(&n)

	ans := 0
	for i := 0; i <= n; i++ {
		idx[i] = -1
	}
	for i := n + 1; i <= 1600; i++ {
		idx[i] = -i
	}

	if n <= 5 {
		for i := 1; i <= n; i++ {
			ans = max(ans, ask(i))
		}
	} else {
		cl, cr, dl, dr := 0, 1597, 610, 987
		el := ask(dl)
		er := ask(dr)
		ans = max(ans, el, er)
		if el < er {
			cl = dl
			dl = dr
			dr = -1
			el = er
			er = -1
		} else {
			cr = dr
			dr = dl
			dl = -1
			er = el
			el = -1
		}

		for i := 12; i >= 0; i-- {
			if dl == -1 {
				dl = cl + fib[i]
				el = ask(dl)
			} else if dr == -1 {
				dr = cr - fib[i]
				er = ask(dr)
			}
			ans = max(ans, el, er)
			if el < er {
				cl = dl
				dl = dr
				dr = -1
				el = er
				er = -1
			} else {
				cr = dr
				dr = dl
				dl = -1
				er = el
				el = -1
			}
		}
		for i := cl + 1; i <= cr-1; i++ {
			ans = max(ans, ask(i))
		}
	}

	fmt.Printf("! %v\n", ans)

}
