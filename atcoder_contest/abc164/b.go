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
	var d int
	fmt.Fscan(reader, &d)

	for {
		c -= b
		if c <= 0 {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		}
		a -= d
		if a <= 0 {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
	}

}
