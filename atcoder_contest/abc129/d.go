package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}
func lowerBound(value, boader int) bool {
	return boader <= value
}

func binarySearch(a []int, boader int, criteria func(value, boader int) bool) int {
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

func main() {
	defer writer.Flush()
	var h int
	fmt.Fscan(reader, &h)
	var w int
	fmt.Fscan(reader, &w)

	s := make([][]string, h)
	for i := 0; i < h; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		s[i] = strings.Split(tmp, "")
	}

	var a = make([][]int, h)
	var b = make([][]int, w)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == "#" {
				a[i] = append(a[i], j)
			}
		}
		a[i] = append(a[i], w)
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if s[j][i] == "#" {
				b[i] = append(b[i], j)
			}
		}
		b[i] = append(b[i], h)
	}

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == "#" {
				continue
			}
			cnt := 0
			idx := binarySearch(a[i], j, lowerBound)
			if 0 <= idx-1 && idx < len(a[i]) {
				cnt += a[i][idx] - a[i][idx-1] - 1
			} else if idx < len(a[i]) {
				cnt += a[i][idx]
			}

			idx = binarySearch(b[j], i, lowerBound)
			if 0 <= idx-1 && idx < len(b[j]) {
				cnt += b[j][idx] - b[j][idx-1] - 1
			} else if idx < len(b[j]) {
				cnt += b[j][idx]
			}

			// fmt.Fprintf(writer, "%v %v %v\n", i, j, cnt)
			ans = max(cnt-1, ans)
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
