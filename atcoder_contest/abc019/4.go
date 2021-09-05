package main

import (
	"fmt"
	"sort"
)

type pair struct {
	first  int
	second int
}

func main() {
	var n int
	fmt.Scan(&n)

	ans := []pair{}
	for i := 2; i <= n; i++ {
		fmt.Printf("? 1 %v\n", i)
		var c int
		fmt.Scan(&c)
		ans = append(ans, pair{first: i, second: c})
	}
	sort.Slice(ans, func(i, j int) bool { return ans[i].second > ans[j].second })
	start := ans[0].first

	ans = []pair{}
	for i := 1; i <= n; i++ {
		if i == start {
			continue
		}
		fmt.Printf("? %v %v\n", start, i)
		var c int
		fmt.Scan(&c)
		ans = append(ans, pair{first: i, second: c})
	}
	sort.Slice(ans, func(i, j int) bool { return ans[i].second > ans[j].second })

	fmt.Printf("! %v\n", ans[0].second)
}
