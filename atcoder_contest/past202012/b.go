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
	var n int
	fmt.Fscan(reader, &n)
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	t := ""
	used := map[string]int{}
	for i := 0; i < len(s); i++ {
		c := s[len(s)-i-1]
		if used[c] == 0 {
			t += c
			used[c]++
		}
	}

	ans := []rune(t)
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}
	fmt.Fprintf(writer, "%v\n", string(ans))
}
