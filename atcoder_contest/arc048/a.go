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

	if a*b < 0 {
		fmt.Fprintf(writer, "%v\n", abs(b-a)-1)
	} else {
		fmt.Fprintf(writer, "%v\n", abs(b-a))
	}

}
