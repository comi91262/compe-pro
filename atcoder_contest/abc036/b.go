package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	a := make([][]string, n)
	for i := 0; i < n; i++ {
		var ss string
		fmt.Fscan(reader, &ss)
		var s = strings.Split(ss, "")

		a[i] = s
	}

	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			a[i][j], a[n-1-i][j] = a[n-1-i][j], a[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i == j {
				continue
			}

			a[i][j], a[j][i] = a[j][i], a[i][j]

		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%v\n", strings.Join(a[i], ""))
	}

}
