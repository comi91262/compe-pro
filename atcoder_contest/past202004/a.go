package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func convert(s string) (n int) {
	switch {
	case s[0] == "B"[0]:
		n = 9 - int(s[1]-"0"[0])
	case s[1] == "F"[0]:
		n = int(s[0]-"0"[0]) + 8
	}

	return
}
func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func main() {
	defer writer.Flush()
	var s string
	fmt.Fscan(reader, &s)
	var t string
	fmt.Fscan(reader, &t)

	fmt.Fprintf(writer, "%v\n", abs(convert(t)-convert(s)))
}
