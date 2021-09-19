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

	var n int
	fmt.Fscan(reader, &n)
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")
	l := 0
	for i := 0; i < len(s); i++ {
		if s[i] == "#" {
			l = i
			break
		}
	}
	r := 0
	for i := 0; i < len(s); i++ {
		if s[len(s)-i-1] == "#" {
			r = i
			break
		}
	}
	mx := 0
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == "#" {
			mx = max(mx, cnt)
			cnt = 0
			continue
		}
		cnt++
	}

	if mx <= l+r {
		fmt.Fprintf(writer, "%v %v\n", l, r)
	} else {
		fmt.Fprintf(writer, "%v %v\n", l+mx-l-r, r)
	}
}
