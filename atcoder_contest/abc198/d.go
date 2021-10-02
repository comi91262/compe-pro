package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func dfsD(size, depth, lower, upper int, memo map[int]int, a []int, accum *[][]int) {
	if size == depth {
		b := make([]int, len(a))
		copy(b, a)
		*accum = append(*accum, b)
		return
	}

	for i := lower; i <= upper; i++ {
		if memo[i] > 0 {
			continue
		}
		a = append(a, i)
		memo[i]++
		dfsD(size, depth+1, lower, upper, memo, a, accum)
		a = a[:len(a)-1]
		memo[i]--
	}
}

func main() {
	defer writer.Flush()
	var ss string
	fmt.Fscan(reader, &ss)
	var s1 = strings.Split(ss, "")
	fmt.Fscan(reader, &ss)
	var s2 = strings.Split(ss, "")
	fmt.Fscan(reader, &ss)
	var s3 = strings.Split(ss, "")

	m := map[string]int{}
	for i := range s1 {
		m[s1[i]]++
	}
	for i := range s2 {
		m[s2[i]]++
	}
	for i := range s3 {
		m[s3[i]]++
	}
	if len(m) > 10 {
		fmt.Fprintf(writer, "%v\n", "UNSOLVABLE")
		return
	}
	ch := []string{}
	for k := range m {
		ch = append(ch, k)
	}

	cand := [][]int{}
	dfsD(len(m), 0, 0, 9, map[int]int{}, []int{}, &cand)

	for _, a := range cand {
		for i, c := range ch {
			m[c] = a[i]
		}
		if m[s1[0]] == 0 || m[s2[0]] == 0 || m[s3[0]] == 0 {
			continue
		}

		x := 0
		for i := range s1 {
			x *= 10
			x += m[s1[i]]
		}
		y := 0
		for i := range s2 {
			y *= 10
			y += m[s2[i]]
		}
		z := 0
		for i := range s3 {
			z *= 10
			z += m[s3[i]]
		}
		if x > 0 && y > 0 && z > 0 && x+y == z {
			fmt.Fprintf(writer, "%v\n%v\n%v\n", x, y, z)
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", "UNSOLVABLE")
}
