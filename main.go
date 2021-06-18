package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	fmt.Fprintf(writer, "%d\n", n)
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
// const d = 1_000_000_000 + 7
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
