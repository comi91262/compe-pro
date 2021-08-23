package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	sum := 0
	pre := 0
	for i := 0; i < n; i++ {
		sum += abs(a[i] - pre)
		pre = a[i]
	}
	sum += abs(0 - pre)

	pre = 0
	for i := 0; i < n-1; i++ {
		fmt.Fprintf(writer, "%v\n", sum-abs(a[i]-pre)-abs(a[i]-a[i+1])+abs(pre-a[i+1]))
		pre = a[i]
	}
	fmt.Fprintf(writer, "%v\n", sum-abs(a[n-1]-pre)-abs(a[n-1]-0)+abs(pre-0))

}
