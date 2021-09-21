package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func dfs(n, d int, a string, accum *[]string) {
	if 0 < d && d <= n {
		*accum = append(*accum, a)
	}
	if d == n {
		return
	}

	for i := 0; i < 26; i++ {
		a += string("a"[0] + byte(i))
		dfs(n, d+1, a, accum)
		a = a[:len(a)-1]
	}
	a += "."
	dfs(n, d+1, a, accum)
	a = a[:len(a)-1]
}

func check(s, t string) bool {
	for i := 0; i < len(s)-len(t)+1; i++ {
		isOk := true
		for j := 0; j < len(t); j++ {
			if t[j] == "."[0] {
				continue
			}
			if s[i+j] != t[j] {
				isOk = false
			}
		}
		if isOk {
			return true
		}
	}
	return false
}

func main() {
	defer writer.Flush()

	var s string
	fmt.Fscan(reader, &s)
	accum := []string{}
	dfs(3, 0, "", &accum)

	ans := 0
	for _, t := range accum {
		if check(s, t) {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
