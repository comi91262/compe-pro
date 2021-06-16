package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	var a = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	var b = make([]int, n+1)
	for i := 1; i < n; i++ {
		b[i] = a[i+1] - a[i]
	}

	sum := 0
	for i := 1; i < n; i++ {
		sum += abs(b[i])
	}

	var l = make([]int, q+1)
	var r = make([]int, q+1)
	var v = make([]int, q+1)
	for i := 1; i < q+1; i++ {
		fmt.Fscan(reader, &l[i], &r[i], &v[i])
	}

	for i := 1; i < q+1; i++ {
		mae := abs(b[l[i]-1]) + abs(b[r[i]])
		if l[i] >= 2 {
			b[l[i]-1] += v[i]
		}
		if r[i] <= n-1 {
			b[r[i]] -= v[i]
		}
		ato := abs(b[l[i]-1]) + abs(b[r[i]])
		sum += (ato - mae)
		fmt.Fprintf(writer, "%d\n", sum)
	}

}

//
// var n int
// fmt.Fscan(reader, &n)
// fmt.Fprintf(writer, "%d\n", n)

//	var n int
//	fmt.Fscan(reader, &n)
//	var a = make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Fscan(reader, &a[i])
//	}
// const MAX = 1_000_000_000 + 7
// for i := 0; i < n; i++ {
// for j := 0; j < n; j++ {
// }
// }
// func abs(x int) int {
// 	return int(math.Abs(float64(x)))
// }
// func min(x, y int) int {
// 	return int(math.Min(float64(x), float64(y)))
// }
// func max(x, y int) int {
// 	return int(math.Max(float64(x), float64(y)))
// }

//	var n, k int
//	fmt.Fscan(reader, &n, &k)
//
//	var s = make([]sum, n)
//	for i := 0; i < n; i++ {
//		fmt.Fscan(reader, &s[i].a, &s[i].b)
//	}
//	sort.Slice(s, func(i, j int) bool { return s[i].a > s[j].a })
//
//	ans := 0
//	for i := 0; i < n; i++ {
//		if k == 0 {
//			break
//		}
//		if k == 1 || k == 2 {
//			s2 := make([]int, n-i)
//			for j := 0; j < n-i; j++ {
//				s2[j] = s[j+i].b
//			}
//			sort.Slice(s2, func(i, j int) bool { return s2[i] > s2[j] })
//			fmt.Fprintf(writer, "%v\n", s)
//			fmt.Fprintf(writer, "%v\n", s2)
//
//			if k == 2 {
//				ans += s2[1] + s2[0]
//				break
//			}
//			if k == 1 {
//				ans += s2[0]
//				break
//			}
//		}
//
//		ans += s[i].a
//		fmt.Fprintf(writer, "%v\n", s[i].a)
//		k -= 2
//	}
