package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const inf = 300001

func upperBound(value, boader int) bool {
	return boader < value
}

func lowerBound(value, boader int) bool {
	return boader <= value
}

func binarySearch(a []int, boader int, criteria func(value, boader int) bool) int {
	abs := func(x int) int {
		if x >= 0 {
			return x
		} else {
			return x * -1
		}
	}

	ng := -1
	ok := len(a)

	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		if criteria(a[mid], boader) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

func lis(a []int) []int {
	r := make([]int, len(a))

	b := make([]int, 0)
	for i := 0; i < len(a); i++ {
		cnt := binarySearch(b, a[i], lowerBound)

		if cnt == len(b) {
			b = append(b, a[i])
		} else {
			b[cnt] = a[i]
		}
		r[i] = cnt + 1
	}

	return r
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	var b = make([]int, n)
	var t int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t)
		a[i] = t
		b[n-1-i] = t
	}

	p := lis(a)
	q := lis(b)

	mx := 0
	for i := 0; i < n; i++ {
		mx = max(p[i]+q[n-1-i]-1, mx)
	}

	fmt.Fprintf(writer, "%v\n", mx)

}

// var n, q int
// fmt.Fscan(reader, &n, &q)

// var s string
// fmt.Fscan(reader, &n, &k, &s)
// ss := strings.Split(s, "")

// var a = make([]int, n)
// for i := 0; i < n; i++ {
// 	fmt.Fscan(reader, &a[i])
// }
// var n int
// fmt.Fscan(reader, &n)
// fmt.Fprintf(writer, "%d\n", n)

//	var n int
//	fmt.Fscan(reader, &n)
//	var a = make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Fscan(reader, &a[i])
//	}
// const d = 1_000_000_000 + 7
// for i := 0; i < n; i++ {
// for j := 0; j < n; j++ {
// for k := 0; k < n; k++ {
// }
// }
// }

// var g = make([][]int, n+1)
// var a = make([]int, m+1)
// var b = make([]int, m+1)
// for i := 1; i < m+1; i++ {
// 	fmt.Fscan(reader, &a[i], &b[i])
// 	g[a[i]] = append(g[a[i]], b[i])
// 	g[b[i]] = append(g[b[i]], a[i])
// }

// m := map[int]int{}
// a := make([]int, n)
// for i := 0; i < len(a); i++ {
// 	m[a[i]]++
// }

// func abs(x int) int {
// 	if x >= 0 {
// 		return x
// 	} else {
// 		return x * -1
// 	}
// }
//
// func min(arg ...int) int {
// 	min := arg[0]
// 	for _, x := range arg {
// 		if min > x {
// 			min = x
// 		}
// 	}
// 	return min
// }
//
func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}
