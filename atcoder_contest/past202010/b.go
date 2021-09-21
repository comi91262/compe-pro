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

	var x int
	fmt.Fscan(reader, &x)
	var y int
	fmt.Fscan(reader, &y)

	if y == 0 {
		fmt.Fprintf(writer, "%v\n", "ERROR")
		return
	}

	ans := x * 100 / y
	fmt.Fprintf(writer, "%v.%02v\n", ans/100, ans%100)
}
