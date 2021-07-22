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
	var s string
	fmt.Fscan(reader, &n, &s)

	for i := range s {
		if (s[i] - 48) == 1 {
			if i%2 == 0 {
				fmt.Fprintf(writer, "%v\n", "Takahashi")
			} else {
				fmt.Fprintf(writer, "%v\n", "Aoki")
			}
			return
		}
	}

}

