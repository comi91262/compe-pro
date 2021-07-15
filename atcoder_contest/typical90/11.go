package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

type task struct {
	d int
	c int
	s int
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	t := make([]task, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t[i].d, &t[i].c, &t[i].s)
	}

	sort.Slice(t, func(i, j int) bool { return t[i].d < t[j].d })

	days := 5001
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, days)
	}

	for i := 0; i < n; i++ {
		for day := 1; day < days; day++ {
			if t[i].c <= day && day <= t[i].d {
				dp[i+1][day] = max(dp[i][day-t[i].c]+t[i].s, dp[i][day])
			} else {
				dp[i+1][day] = dp[i][day]
			}
		}
	}

	// mx := 0
	// for i := 0; i < 1<<n; i++ {
	// 	set := []int{}
	// 	for j := 0; j < n; j++ {
	// 		if i&(1<<j) != 0 {
	// 			set = append(set, j)
	// 		}
	// 	}

	// 	// fmt.Fprintf(writer, "%v\n", set)
	// 	day := 1
	// 	bill := 0
	// 	for _, idx := range set {
	// 		if day >= t[idx].d || day+t[idx].c > t[idx].d {
	// 			continue
	// 		}
	// 		day += t[idx].c
	// 		bill += t[idx].s
	// 	}

	mx := 0
	for i := 0; i < days; i++ {
		mx = max(mx, dp[n][i])
	}

	fmt.Fprintf(writer, "%v\n", mx)
}
