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
	fmt.Fscan(reader, &ss)
	var t = strings.Split(ss, "")

	ans := 1 << 60
	for i := 0; i < len(s)-len(t)+1; i++ {
		cnt := 0
		for j := 0; j < len(t); j++ {
			if s[i+j] != t[j] {
				cnt++
			}
		}
		ans = min(ans, cnt)
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
