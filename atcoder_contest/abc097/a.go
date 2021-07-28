package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}

}

func main() {
	defer writer.Flush()

	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)
	var d int
	fmt.Fscan(reader, &d)

	if abs(a-c) <= d {
		fmt.Fprintf(writer, "%v\n", "Yes")
	} else if abs(a-b) <= d && abs(b-c) <= d {
		fmt.Fprintf(writer, "%v\n", "Yes")
	} else {
		fmt.Fprintf(writer, "%v\n", "No")
	}
}
