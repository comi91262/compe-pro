package main

import (
	"bufio"
	"fmt"
	"os"
)

// type info struct {
// 	start int
// 	w     int
// }
//
// func abs(x, y int) int {
// 	fx := float64(x)
// 	fy := float64(y)
// 	return int(math.Abs(fx - fy))
// }

var deque [20000000]int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var q int
	fmt.Fscan(reader, &q)

	var t = make([]int, q)
	var x = make([]int, q)

	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &t[i], &x[i])
	}

	var cl, cr = 10000000, 10000000
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			deque[cl-1] = x[i]
			cl--
		case 2:
			deque[cr] = x[i]
			cr++
		case 3:
			fmt.Fprintf(writer, "%d\n", deque[cl+x[i]-1])
		}
	}
}
