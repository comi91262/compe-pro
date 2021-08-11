package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type sword struct {
	s int
	t int
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var h int
	fmt.Fscan(reader, &h)

	var k = make([]sword, n)
	var a = make([]int, n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		fmt.Fscan(reader, &b[i])
		k[i] = sword{a[i], b[i]}
	}

	sort.Slice(k, func(i, j int) bool { return k[i].s > k[j].s })

	var t = []sword{}
	var mx = k[0].s
	for i := 0; i < n; i++ {
		if mx > k[i].t {
			continue
		}
		t = append(t, k[i])
	}

	sort.Slice(t, func(i, j int) bool { return t[i].t > t[j].t })

	ans := 0
	sum := 0
	for i := range t {
		sum += t[i].t
		ans++
		if h <= sum {
			fmt.Fprintf(writer, "%v\n", ans)
			return
		}
	}

	sort.Slice(t, func(i, j int) bool { return t[i].s > t[j].s })
	for {
		sum += t[0].s
		ans++
		if h <= sum {
			fmt.Fprintf(writer, "%v\n", ans)
			return
		}
	}
}
