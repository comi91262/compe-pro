package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)
	s := make([][]string, n)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		s[i] = strings.Split(tmp, "")
	}
	dx := [9]int{1, 1, 1, 0, 0, 0, -1, -1, -1}
	dy := [9]int{1, 0, -1, 1, 0, -1, 1, 0, -1}

	var a = make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			cnt := 0
			for k := 0; k < 9; k++ {
				nx := i + dx[k]
				ny := j + dy[k]
				if 0 <= nx && nx < n && 0 <= ny && ny < m {
					if s[nx][ny] == "#" {
						cnt++
					}
				}
			}
			a[i][j] = cnt
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fprintf(writer, "%v", a[i][j])
		}
		fmt.Fprintf(writer, "\n")
	}

}
