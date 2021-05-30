package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func min(x, y, z int) int {
	var m = int(math.Min(float64(x), float64(y)))
	return int(math.Min(float64(m), float64(z)))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func div(a, b int) int {
	cnt := 1
	for a-b > 0 {
		a = a - b
		cnt++
	}
	return cnt
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b, c int
	fmt.Fscan(reader, &a, &b, &c)

	var n, m, l = gcd(a, b), gcd(b, c), gcd(c, a)
	var u = min(n, m, l)

	fmt.Fprintf(writer, "%d\n", a/u-1+b/u-1+c/u-1)
	//	for i := 0; i < q; i++ {
	//	}
	//		fmt.Fprintf(writer, "%d\n", v)
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
