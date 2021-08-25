package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}
func main() {
	defer writer.Flush()
	var n string
	fmt.Fscan(reader, &n)

	var s = strings.Split(n, "")

	a := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		a[i], _ = strconv.Atoi(s[i])
	}

	m := map[int]int{}
	for i := 0; i < len(a); i++ {
		m[a[i]%3]++
	}

	ans := abs(m[1]-m[2]) % 3

	if m[1]%3 == 0 && m[2]%3 == 0 {
		ans = 0
	} else if m[1]%3 == 0 {
		ans = min(ans, m[2]%3)
	} else if m[2]%3 == 0 {
		ans = min(ans, m[1]%3)
	}

	if ans >= len(a) {
		ans = -1
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
