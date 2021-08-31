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

func main() {
	defer writer.Flush()

	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	if len(s) == 1 {
		t := strings.Join(s, "")
		tt, _ := strconv.Atoi(t)
		if tt%8 == 0 {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		} else {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
	}

	m := map[int]int{}
	for i := 0; i < len(s); i++ {
		n, _ := strconv.Atoi(s[i])
		m[n]++
	}

	if len(s) == 2 {
		check := []int{}
		for i := 10; i < 100; i++ {
			if i%8 == 0 {
				check = append(check, i)
			}
		}

		for i := range check {
			a := check[i] % 10
			b := check[i] / 10 % 10
			if a == 0 || b == 0 {
				continue
			}
			l := map[int]int{}
			l[a]++
			l[b]++

			if l[a] <= m[a] && l[b] <= m[b] {
				fmt.Fprintf(writer, "%v\n", "Yes")
				return
			}
		}
		fmt.Fprintf(writer, "%v\n", "No")
		return
	}

	check := []int{}
	for i := 100; i < 1000; i++ {
		if i%8 == 0 {
			check = append(check, i)
		}
	}

	for i := range check {
		a := check[i] % 10
		b := check[i] / 10 % 10
		c := check[i] / 10 / 10 % 10
		if a == 0 || b == 0 || c == 0 {
			continue
		}
		l := map[int]int{}
		l[a]++
		l[b]++
		l[c]++

		if l[a] <= m[a] && l[b] <= m[b] && l[c] <= m[c] {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		}

	}
	fmt.Fprintf(writer, "%v\n", "No")
}
