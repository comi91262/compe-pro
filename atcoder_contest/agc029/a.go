package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	a := make([]int, len(s))
	for i := 0; i < len(s)-1; i++ {
		if s[i] == "B" {
			a[i]++
		}
	}
	for i := 0; i < len(s)-1; i++ {
		a[i+1] += a[i]
	}

	ans := 0
	for i := 1; i < len(s); i++ {
		if s[i] == "W" {
			ans += a[i-1]
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
