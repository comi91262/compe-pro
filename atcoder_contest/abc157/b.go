package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func contains(x int, a []int) bool {
	for _, e := range a {
		if x == e {
			return true
		}
	}
	return false
}

func main() {
	defer writer.Flush()

	a := make([][]int, 3)
	for i := 0; i < 3; i++ {
		a[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	var n int
	fmt.Fscan(reader, &n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if contains(a[i][j], b) {
				a[i][j] = 0
			}
		}
	}

	for i := 0; i < 3; i++ {
		if a[i][0] == 0 && a[i][1] == 0 && a[i][2] == 0 {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		}
	}

	for i := 0; i < 3; i++ {
		if a[0][i] == 0 && a[1][i] == 0 && a[2][i] == 0 {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		}
	}

	if a[0][0] == 0 && a[1][1] == 0 && a[2][2] == 0 {
		fmt.Fprintf(writer, "%v\n", "Yes")
		return
	}

	if a[0][2] == 0 && a[1][1] == 0 && a[2][0] == 0 {
		fmt.Fprintf(writer, "%v\n", "Yes")
		return
	}

	fmt.Fprintf(writer, "%v\n", "No")
}
