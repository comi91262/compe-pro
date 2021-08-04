package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func main() {
	defer writer.Flush()

	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	zn := 0
	on := 0
	for i := range s {
		switch s[i] {
		case "0":
			zn++
		case "1":
			on++
		}
	}
	fmt.Fprintf(writer, "%v\n", 2*min(zn, on))

}
