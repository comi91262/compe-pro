package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	m := make([]map[string]int, n)
	a := make([][]string, n)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		a[i] = strings.Split(tmp, "")
	}
	for i := 0; i < n; i++ {
		m[i] = map[string]int{}
		for j := 0; j < n; j++ {
			m[i][a[i][j]]++
		}
	}
	st := ""
	for i := 0; i < n/2; i++ {
		isOk := false
		for j := 0; j < 26; j++ {
			c := string("a"[0] + byte(j))
			s := m[i][c]
			t := m[n-i-1][c]
			if s > 0 && t > 0 {
				isOk = true
				st += c
			}
		}
		if !isOk {
			fmt.Fprintf(writer, "%v\n", -1)
			return
		}
	}

	mid := ""
	if n%2 == 1 {
		mid = a[n/2][0]
	}
	fmt.Fprintf(writer, "%v\n", st+mid+reverseString(st))
}
