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
	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	if n == 1 {
		fmt.Fprintf(writer, "%v\n", 1)
		return
	}

	ans := 0
	so := 0
	for i := 0; i < n-1; i++ {
		switch {
		case so == 1 && a[i+1] > a[i]:
		case so == 1 && a[i+1] < a[i]:
			ans++
			so = 0
		case so == 0 && a[i+1] > a[i]:
			so = 1
		case so == 0 && a[i+1] < a[i]:
			so = -1
		case so == 0 && a[i+1] == a[i]:
		case so == -1 && a[i+1] > a[i]:
			ans++
			so = 0
		case so == -1 && a[i+1] == a[i]:
		case so == -1 && a[i+1] < a[i]:
		}
	}
	fmt.Fprintf(writer, "%v\n", ans+1)
}
