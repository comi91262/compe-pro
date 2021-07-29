package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var s string
	fmt.Fscan(reader, &s)

	mx := 0
	cnt := 0
	for i := range s {
		switch s[i] {
		case "I"[0]:
			cnt++
		case "D"[0]:
			cnt--
		}
		mx = max(mx, cnt)
	}

	fmt.Fprintf(writer, "%v\n", mx)
}
