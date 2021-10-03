package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type pair struct {
	first  int
	second int
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var t int
	fmt.Fscan(reader, &t)

	var a = make([]int, n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
	}
	var p = make([]pair, n)
	for i := 0; i < n; i++ {
		p[i].first = a[i]
		p[i].second = b[i]
	}

	tot := 0
	for i := 0; i < n; i++ {
		tot += p[i].second
	}
	if tot > t {
		fmt.Fprintf(writer, "%v\n", -1)
		return
	}

	sort.Slice(p, func(i, j int) bool { return p[i].first-p[i].second < p[j].first-p[j].second })

	ans := 0
	for i := 0; i < n; i++ {
		if tot+p[i].first-p[i].second > t {
			break
		}
		tot += p[i].first - p[i].second
		ans++
	}
	fmt.Fprintf(writer, "%v\n", n-ans)

}
