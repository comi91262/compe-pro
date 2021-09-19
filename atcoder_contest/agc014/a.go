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
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)

	if a%2 == 1 || b%2 == 1 || c%2 == 1 {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}

	if a == b && b == c && c == a {
		fmt.Fprintf(writer, "%v\n", -1)
		return
	}
	ans := 0
	for {
		if a%2 == 1 || b%2 == 1 || c%2 == 1 {
			break
		}
		ans++
		a, b, c = (b+c)/2, (c+a)/2, (a+b)/2
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
