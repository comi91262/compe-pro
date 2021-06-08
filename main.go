package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	fmt.Fprintf(writer, "%d\n", sum)
}

// var n int
// fmt.Fscan(reader, &n)
// fmt.Fprintf(writer, "%d\n", sum)

// for i := 1; i < n; i++ {
// 	for j := 1; j < n; j++ {
// 	}
// }

// var n int
//   var a = make([]int, n)
// for i := 0; i < n; i++ {
// 	fmt.Fscan(reader, &a[i])
// }
