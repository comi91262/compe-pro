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
	for i := 0; i < n-1; i++ {
		switch {
		case a[i+1] > a[i]:
			fmt.Fprintf(writer, "up %v\n", a[i+1]-a[i])
		case a[i+1] == a[i]:
			fmt.Fprintf(writer, "stay \n")
		case a[i+1] < a[i]:
			fmt.Fprintf(writer, "down %v\n", a[i]-a[i+1])
		}

	}
}
