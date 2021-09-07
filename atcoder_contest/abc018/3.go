package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type pair struct {
	first  int
	second int
}

func main() {
	defer writer.Flush()

	var r int
	fmt.Fscan(reader, &r)
	var c int
	fmt.Fscan(reader, &c)
	var k int
	fmt.Fscan(reader, &k)

	s := make([][]string, r)
	for i := 0; i < r; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		s[i] = strings.Split(tmp, "")
	}

	var p = make([][]pair, r)
	for i := 0; i < r; i++ {
		p[i] = make([]pair, c)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if s[i][j] == "o" {
				p[i][j].first = 1
				p[i][j].second = 1
			}
		}
	}
	for i := 0; i < c; i++ {
		for j := 0; j < r-1; j++ {
			if p[j+1][i].first != 0 {
				p[j+1][i].first += p[j][i].first
			}
			if p[r-1-j-1][i].second != 0 {
				p[r-1-j-1][i].second += p[r-1-j][i].second
			}
		}
	}

	ans := 0
	for i := k - 1; i <= r-k; i++ {
		for j := k - 1; j <= c-k; j++ {
			flag := true
			for l := 1 - k; l <= k-1; l++ {
				t := l
				if l < 0 {
					t *= -1
				}
				if p[i][j+l].first < k-t || p[i][j+l].second < k-t {
					flag = false
					break
				}
			}
			if flag {
				ans++
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
