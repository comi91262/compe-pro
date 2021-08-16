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
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	var s = make([]string, n)
	var d = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
		fmt.Fscan(reader, &d[i])
	}

	x := 0
	for i := 0; i < n; i++ {
		tmp := 1
		switch s[i] {
		case "West":
			tmp *= -1
		case "East":
			tmp *= 1
		}

		switch {
		case d[i] < a:
			x += a * tmp
		case a <= d[i] && d[i] <= b:
			x += d[i] * tmp
		case b < d[i]:
			x += b * tmp
		}
	}

	switch {
	case x < 0:
		fmt.Fprintf(writer, "West %v\n", x*-1)
	case 0 < x:
		fmt.Fprintf(writer, "East %v\n", x)
	case x == 0:
		fmt.Fprintf(writer, "%v\n", 0)
	}

}
