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
	var mx = make([]float64, n)
	var mn = make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &mx[i])
		fmt.Fscan(reader, &mn[i])
	}

	ans := make([]int, 6)
	for i := 0; i < n; i++ {
		switch {
		case mx[i] >= 35.0:
			ans[0]++
		case 30.0 <= mx[i] && mx[i] < 35.0:
			ans[1]++
		case 25.0 <= mx[i] && mx[i] < 30.0:
			ans[2]++
		}

		switch {
		case mn[i] >= 25.0:
			ans[3]++
		case mn[i] < 0.0 && 0.0 <= mx[i]:
			ans[4]++
		case mx[i] < 0.0:
			ans[5]++
		}
	}
	for i := 0; i < len(ans); i++ {
		if i > 0 {
			fmt.Fprintf(writer, " ")
		}
		fmt.Fprintf(writer, "%v", ans[i])
	}
	fmt.Fprintf(writer, "\n")
}
