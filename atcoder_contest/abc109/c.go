package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var x int
	fmt.Fscan(reader, &x)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	a = append(a, x)
	sort.Ints(a)

	//if len(a) == 1 {
	//	fmt.Fprintf(writer, "%v\n", abs(a[0]-x))
	//	return
	//}

	var b = make([]int, n+1)
	for i := 0; i < n; i++ {
		b[i] = a[i+1] - a[i]
	}

	ans := 0
	for i := range b {
		ans = gcd(ans, b[i])
	}

	fmt.Fprintf(writer, "%v\n", ans)
}

