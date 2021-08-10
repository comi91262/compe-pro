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
	var s string
	fmt.Fscan(reader, &s)

	cand := "b"
	if s == cand {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}

	for i := 1; i <= 100; i++ {
		switch i % 3 {
		case 0:
			cand = "b" + cand + "b"
		case 1:
			cand = "a" + cand + "c"
		case 2:
			cand = "c" + cand + "a"
		}

        if s == cand {
            fmt.Fprintf(writer, "%v\n", i)
            return
        }
	}

    fmt.Fprintf(writer, "%v\n", -1)
}
