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
	var b = make([]int, n+1)
	for i := 0; i < n; i++ {
		b[a[i]]++
	}

	for i := 0; i < n+1; i++ {
		if b[i] > 1 {
			fmt.Fprintf(writer, "No\n")
			return
		}
	}
	fmt.Fprintf(writer, "Yes\n")
}
