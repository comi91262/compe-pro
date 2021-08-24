package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func combination(n, k int) int {
	r := 1
	for d := 1; d <= k; d++ {
		r *= n
		n--
		r /= d

	}

	return r
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func pow(a, x int) int {
	r := 1
	for x > 0 {
		if x&1 == 1 {
			r *= a
		}
		a *= a
		x >>= 1
	}
	return r
}

func main() {
	defer writer.Flush()

	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	m := map[string]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	// fmt.Fprintf(writer, "%v\n", m)
	a := m["o"]
	b := m["?"]
	if a+b == 0 {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}

	ans := 0
	c := combination

	switch a {
	case 0:
		ans = pow(b, 4)
	case 1:
		ans += c(4, 1) * pow(b, 3)
		ans += c(4, 2) * pow(b, 2)
		ans += c(4, 3) * pow(b, 1)
		ans += c(4, 4) * pow(b, 0)
	case 2:
		ans += c(4, 2) * 2 * pow(b, 2)
		ans += c(4, 3) * c(3, 1) * 2 * pow(b, 1)
		ans += c(4, 4) * (c(4, 1)*2 + c(4, 2)) * pow(b, 0)
	case 3:
		ans += c(4, 3) * fact(3) * pow(b, 1)
		ans += c(4, 4) * c(4, 2) * fact(3) * pow(b, 0)
	case 4:
		ans += c(4, 4) * c(a, 4) * fact(4) * pow(b, 0)
	default:
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
