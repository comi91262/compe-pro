package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	first  int
	second int
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var p = make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i].first, &p[i].second)
	}

	at := 0
	bt := 0
	for i := 0; i < n; i++ {
		at += p[i].first
	}

	if at < bt {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}

	sort.Slice(p, func(i, j int) bool {
		s := p[i].first*2 + p[i].second
		t := p[j].first*2 + p[j].second

		return s > t
	})

	cnt := 0
	for i := 0; i < n; i++ {
		at -= p[i].first
		bt += p[i].first + p[i].second
		cnt++
		if at < bt {
			fmt.Fprintf(writer, "%v\n", cnt)
			return
		}

	}
}
