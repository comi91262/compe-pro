package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func toDigits(x, base int) []int {
	if x == 0 {
		return []int{0}
	}

	ans := make([]int, 0)
	for x != 0 {
		ans = append(ans, x%base)
		x = x / base
	}
	return ans
}

const d = 100000

func main() {
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	var a = make([]int, d)
	m := n
	a[m]++
	period := 0
	for {
		period++
		sum := 0
		for _, v := range toDigits(m, 10) {
			sum += v
		}
		m = (m + sum) % d
		if a[m] > 0 {
			break
		} else {
			a[m]++
		}
	}

	var b = make([]int, d)
	b[m]++
	cycle := 0
	for {
		cycle++
		sum := 0
		for _, v := range toDigits(m, 10) {
			sum += v
		}
		m = (m + sum) % d
		if b[m] > 0 {
			break
		} else {
			b[m]++
		}
	}

	if k > period {
		k = (k-period)%cycle + period
	}

	for k > 0 {
		sum := 0
		for _, v := range toDigits(n, 10) {
			sum += v
		}
		n = (n + sum) % d
		k--
	}
	fmt.Fprintf(writer, "%d\n", n)
}

// var n, q int
// fmt.Fscan(reader, &n, &q)

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
// }
// }
// func max(x, y int) int {
// 	return int(math.Max(float64(x), float64(y)))
// }

// var g = make([][]int, n+1)
// var a = make([]int, n+1)
// var b = make([]int, n+1)
// for i := 1; i < n; i++ {
// 	fmt.Fscan(reader, &a[i], &b[i])
// 	g[a[i]] = append(g[a[i]], b[i])
// 	g[b[i]] = append(g[b[i]], a[i])
// }

// m := map[int]int{}
// a := make([]int, n)
// for i := 0; i < len(a); i++ {
// 	m[a[i]]++
// }
