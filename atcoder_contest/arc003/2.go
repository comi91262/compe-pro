package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type pair struct {
	first  string
	second string
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var p = make([]pair, n)
	var s = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
		p[i].first = s[i]
		p[i].second = reverseString(s[i])
	}
	sort.Slice(p, func(i, j int) bool { return p[i].second < p[j].second })

	for i := range p {
		fmt.Fprintf(writer, "%v\n", p[i].first)
	}
}
