package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	p := 0
	w := 0
	l := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case "p":
			if p == 0 {
				l++
				p++
				continue
			}
			p--
		case "g":
			if p != 0 {
				w++
				p--
				continue
			}
			p++
		}
	}
	// fmt.Fprintf(writer, "%v %v\n", w, l)
	fmt.Fprintf(writer, "%v\n", max(w-l, 0))
}
