package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type point struct {
	x    int
	y    int
	cost int
}

type queue []point

func (q *queue) push(n point) {
	*q = append(*q, n)
}

func (q *queue) pop() point {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

const mod = 1_000_000_000 + 7

func nextInt() int {
	scanner.Scan()
	ret, e := strconv.Atoi(scanner.Text())
	if e != nil {
		panic(e)
	}
	return ret
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	h := nextInt()
	w := nextInt()

	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		for j := 0; j < w; j++ {
			a[i][j] = nextInt()
		}
	}

	var dp = make([][]int, h)
	for i := 0; i < h; i++ {
		dp[i] = make([]int, w)
	}

	var dx = [4]int{1, 0, -1, 0}
	var dy = [4]int{0, 1, 0, -1}

	q := queue{}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			q.push(point{x: i, y: j, cost: a[i][j]})
		}

	}
	sort.Slice(q, func(i, j int) bool { return q[i].cost < q[j].cost })

	for !q.empty() {
		u := q.pop()
		for i := 0; i < 4; i++ {
			tx := u.x + dx[i]
			ty := u.y + dy[i]

			if tx < 0 || tx >= h || ty < 0 || ty >= w {
				continue
			}
			if a[u.x][u.y] >= a[tx][ty] {
				continue
			}
			dp[tx][ty] += dp[u.x][u.y] + 1
			dp[tx][ty] %= mod
		}
	}
	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ans += dp[i][j]
			ans += 1
			ans %= mod
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}

