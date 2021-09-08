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
	var t int
	fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)

		mn := 2 * l
		mx := r

		if mn > mx {
			fmt.Fprintf(writer, "%v\n", 0)
		} else {
			fmt.Fprintf(writer, "%v\n", (mx-mn+1)*(mx-mn+2)/2)
		}

	}

}
