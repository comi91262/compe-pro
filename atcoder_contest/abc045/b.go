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

	var a string
	fmt.Fscan(reader, &a)
	var b string
	fmt.Fscan(reader, &b)
	var c string
	fmt.Fscan(reader, &c)

	turn := "a"
	for {
		switch turn {
		case "a":
			if len(a) == 0 {
				fmt.Fprintf(writer, "%v\n", "A")
				return
			}
			turn = string(a[0])
			a = a[1:]

		case "b":
			if len(b) == 0 {
				fmt.Fprintf(writer, "%v\n", "B")
				return
			}
			turn = string(b[0])
			b = b[1:]
		case "c":
			if len(c) == 0 {
				fmt.Fprintf(writer, "%v\n", "C")
				return
			}

			turn = string(c[0])
			c = c[1:]
		}
	}

}
