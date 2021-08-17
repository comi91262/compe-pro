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

	for a := 0; a <= 25; a++ {
		for b := 0; b <= 14; b++ {
			if a*4+b*7 == n {
				fmt.Fprintf(writer, "%v\n", "Yes")
				return
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", "No")
}
