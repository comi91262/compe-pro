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

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func intersect(x, y map[string]int) map[string]int {
	r := map[string]int{}

	for k, v := range x {
		if y[k] > 0 {
			r[k] = min(v, y[k])
		}
	}

	return r
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	m := map[string]int{}
	for i := 0; i < n; i++ {
		tmp := map[string]int{}
		var ss string
		fmt.Fscan(reader, &ss)
		var s = strings.Split(ss, "")

		for i := range s {
			tmp[s[i]]++
		}
		if i == 0 {
			m = tmp
		} else {
			m = intersect(m, tmp)
		}
	}

	ans := []string{}
	for k, v := range m {
		for i := 0; i < v; i++ {
			ans = append(ans, k)
		}
	}
	sort.Strings(ans)
	fmt.Fprintf(writer, "%v\n", strings.Join(ans, ""))

}
