package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	ans := 0
	for i := a; i <= b; i++ {
		s := strconv.Itoa(i)
		ok := true
		for l := 0; l < len(s); l++ {
			if s[l] != s[len(s)-1-l] {
				ok = false
			}
		}
		if ok {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
