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

	for i := 1; i <= 3; i++ {
		if (a*b*i)%2 == 1 {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "No")
}


