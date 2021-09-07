package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

var b [2][3]int
var c [3][2]int

func calc(v []int) (score int) {
	var table = make([][]int, 3)
	for i := 0; i < 3; i++ {
		table[i] = make([]int, 3)
	}
	for i := 0; i < 9; i++ {
		x := v[i] / 3
		y := v[i] % 3
		if i%2 == 0 {
			table[x][y] = 1
		} else {
			table[x][y] = 0
		}
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if table[i][j] == table[i+1][j] {
				score += b[i][j]
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			if table[i][j] == table[i][j+1] {
				score += c[i][j]
			}
		}
	}

	return
}

func dfs(v []int, visited []bool) (score int) {
	if len(v) == 9 {
		return calc(v)
	}

	mx := -1 << 60
	mn := 1 << 60
	for i := 0; i < 9; i++ {
		if visited[i] {
			continue
		}

		visited[i] = true
		v = append(v, i)
		pre := dfs(v, visited)

		v = v[:len(v)-1]
		visited[i] = false

		if len(v)%2 == 0 {
			if mx < pre {
				mx = pre
				score = pre
			}
		} else {
			if mn > pre {
				mn = pre
				score = pre
			}
		}
	}

	return
}

func main() {
	defer writer.Flush()

	sum := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Fscan(reader, &b[i][j])
			sum += b[i][j]
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Fscan(reader, &c[i][j])
			sum += c[i][j]
		}
	}

	v := []int{}
	visited := make([]bool, 9)
	score := dfs(v, visited)

	fmt.Fprintf(writer, "%v\n", score)
	fmt.Fprintf(writer, "%v\n", sum-score)

}
