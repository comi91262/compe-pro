package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}
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

	var a = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		a[i], _ = strconv.Atoi(s[i])
	}

	mn := 1 << 60
	for i := 0; i < len(a)-2; i++ {
		x, y, z := a[i], a[i+1], a[i+2]
		mn = min(mn, abs(753-100*x-10*y-z))
	}
	fmt.Fprintf(writer, "%v\n", mn)

}
