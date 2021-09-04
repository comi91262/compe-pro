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
	var k int
	fmt.Fscan(reader, &k)

	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")
	fmt.Fscan(reader, &ss)
	var t = strings.Split(ss, "")

	m := map[int]int{}
	for i := 1; i < 10; i++ {
		m[i] += k
	}

	ta := map[int]int{}
	for i := range s {
		if s[i] == "#" {
			continue
		}
		n, _ := strconv.Atoi(s[i])
		ta[n]++
		m[n]--
	}

	ao := map[int]int{}
	for i := range t {
		if t[i] == "#" {
			continue
		}
		n, _ := strconv.Atoi(t[i])
		ao[n]++
		m[n]--
	}
	p := 0
	for _, v := range m {
		p += v
	}
	p = p * (p - 1)

	c := 0
	tat := 0
	for i := 1; i < 10; i++ {
		tat += i * pow(10, ta[i])

	}
	aot := 0
	for i := 1; i < 10; i++ {
		aot += i * pow(10, ao[i])
	}

	cnt := 0
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if m[i] == 0 || m[j] == 0 {
				continue
			}
			cnt++

			tg := i * pow(10, ta[i]) * 9
			ag := j * pow(10, ao[j]) * 9
			if tg+tat > ag+aot {
				if i != j {
					c += m[i] * m[j]
				} else {
					c += m[j] * (m[i] - 1)
				}
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", float64(c)/float64(p))
}
