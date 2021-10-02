package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type pair struct {
	first  string
	second string
}

func main() {
	defer writer.Flush()
	var ss string
	fmt.Fscan(reader, &ss)
	var x = strings.Split(ss, "")

	m := map[string]string{}
	for i := range x {
		m[x[i]] = string(byte(i) + "a"[0])
	}

	var n int
	fmt.Fscan(reader, &n)
	var s = make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i].first)
	}
	for i := 0; i < n; i++ {
		for j := range s[i].first {
			s[i].second += m[string(s[i].first[j])]
		}
	}
	sort.Slice(s, func(i, j int) bool { return s[i].second < s[j].second })
	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%v\n", s[i].first)
	}

}
