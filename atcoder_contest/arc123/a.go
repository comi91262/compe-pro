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

	x := 2*b - a - c

	if x < 0 {
		if x*-1%2 == 0 {
			fmt.Fprintf(writer, "%v\n", x*-1/2)
		} else {
			fmt.Fprintf(writer, "%v\n", (x*-1+1)/2+1)
		}
	} else {
		fmt.Fprintf(writer, "%v\n", x)
	}

}
