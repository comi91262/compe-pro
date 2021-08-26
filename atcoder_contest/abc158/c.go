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

	for x := 0; x <= 1300; x++ {
		if x*8/100 == a && x*10/100 == b {
			fmt.Fprintf(writer, "%v\n", x)
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", -1)
}
