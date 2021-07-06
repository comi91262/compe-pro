package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var a [800][800]int
var min = 1000000001

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}

	mid := k*k/2 + 1
	for i := 0; i < n-k+1; i++ {
		for j := 0; j < n-k+1; j++ {
			b := make([]int, k*k)
			cnt := 0
			for x := i; x < i+k; x++ {
				for y := j; y < j+k; y++ {
					b[cnt] = a[x][y]
					cnt++
				}
			}
			sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
			if min > b[mid-1] {
				min = b[mid-1]
			}
		}
	}

	fmt.Fprintf(writer, "%d\n", min)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

var a, s [4][4]int
var min = 4 //1000_000_000

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	ng := -1
	ok := min
	lim := k*k/2 + 1
	var ext bool
	var mid int

	for (ng + 1) < ok {
		mid = (ng + ok) / 2
		fmt.Fprintf(writer, "%d\n", mid)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				s[i+1][j+1] = s[i+1][j] + s[i][j+1] - s[i][j]
				if a[i][j] > mid {
					s[i+1][j+1]++
				}
			}
		}
		fmt.Fprintf(writer, "%v\n", s)

		ext = false
		for i := 0; i < n-k+1; i++ {
			for j := 0; j < n-k+1; j++ {
				if (s[i+k][j+k] + s[i][j] - s[i][j+k] - s[i+k][j]) < lim {
					ext = true
				}
			}
		}
		if ext {
			ok = mid
		} else {
			ng = mid
		}
	}

	//	for i := 0; i < n-k+1; i++ {
	//		for j := 0; j < n-k+1; j++ {
	//			b := make([]int, k*k)
	//			cnt := 0
	//			for x := i; x < i+k; x++ {
	//				for y := j; y < j+k; y++ {
	//					b[cnt] = a[x][y]
	//					cnt++
	//				}
	//			}
	//			sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
	//			if min > b[mid-1] {
	//				min = b[mid-1]
	//			}
	//		}
	//	}

	fmt.Fprintf(writer, "%d\n", ok)
}

// var n int
// fmt.Fscan(reader, &n)
// fmt.Fprintf(writer, "%d\n", n)

//	var n int
//	fmt.Fscan(reader, &n)
//	var a = make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Fscan(reader, &a[i])
//	}
// snip
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

