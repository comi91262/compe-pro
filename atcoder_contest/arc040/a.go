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
		var tmp string
		fmt.Fscan(reader, &tmp)
		a[i] = strings.Split(tmp, "")
	}

	m := map[string]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m[a[i][j]]++
		}
	}

	switch {
	case m["R"] > m["B"]:
		fmt.Fprintf(writer, "%v\n", "TAKAHASHI")
	case m["R"] == m["B"]:
		fmt.Fprintf(writer, "%v\n", "DRAW")
	case m["R"] < m["B"]:
		fmt.Fprintf(writer, "%v\n", "AOKI")
	}
}
