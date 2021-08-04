package main

import (
	"bufio"
	"fmt"
	"os"
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

type ans struct {
	h int
	x int
	y int
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	var x = make([]int, n)
	var y = make([]int, n)
	var h = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
		fmt.Fscan(reader, &y[i])
		fmt.Fscan(reader, &h[i])
	}

	for i := 0; i < 101; i++ {
		for j := 0; j < 101; j++ {

			m := map[ans]int{}
			for k := 0; k < n; k++ {
				if h[k] == 0 {
					continue
				}
				ax := abs(x[k] - i)
				ay := abs(y[k] - j)
				ah := h[k] + ax + ay

				if ah >= 1 && h[k] == max(ah-ax-ay, 0) {
					m[ans{h: ah, x: i, y: j}]++
				}
			}

			for ans := range m {
				ok := true
				for k := 0; k < n; k++ {
					if h[k] != max(ans.h-abs(x[k]-ans.x)-abs(y[k]-ans.y), 0) {
						ok = false
						break
					}
				}
				if ok {
					fmt.Fprintf(writer, "%v %v %v\n", ans.x, ans.y, ans.h)
					return
				}
			}

		}
	}

}
