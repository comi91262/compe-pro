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

	ans := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			for k := j; k < n; k++ {
				if a[i] == a[j] {
					continue
				}
				if a[j] == a[k] {
					continue
				}
				if a[k] == a[i] {
					continue
				}
				if a[i]+a[j] <= a[k] {
					continue
				}
				if a[i]+a[k] <= a[j] {
					continue
				}
				if a[k]+a[j] <= a[i] {
					continue
				}
				ans++
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
