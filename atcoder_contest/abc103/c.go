package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

//func f(m int, a []int) int {
//	r := 0
//	for i := 0; i < len(a); i++ {
//		r += m % a[i]
//	}
//	return r
//}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans += a[i] - 1
	}
	fmt.Fprintf(writer, "%v\n", ans)
	//for i := 0; i <= 20; i++ {
	//	fmt.Fprintf(writer, "%v %v\n", i, f(i, a))
	//}

}
