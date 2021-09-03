package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func sum(a []int) int {
	r := 0
	for i := range a {
		r += a[i]
	}
	return r
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		cur := 0
		for j := n - i - 1; j < n; j += n - i {
			cur += ans[j]
		}
		switch {
		case a[n-i-1] == 0 && cur%2 == 0:
			ans[n-i-1] = 0
		case a[n-i-1] == 0 && cur%2 == 1:
			ans[n-i-1] = 1
		case a[n-i-1] == 1 && cur%2 == 0:
			ans[n-i-1] = 1
		case a[n-i-1] == 1 && cur%2 == 1:
			ans[n-i-1] = 0
		}
	}

	fmt.Fprintf(writer, "%v\n", sum(ans))
	for i := 0; i < n; i++ {
		if ans[i] == 1 {
			fmt.Fprintf(writer, "%v\n", i+1)
		}
	}

}
