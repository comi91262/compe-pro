package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func main() {
	defer writer.Flush()
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	if len(s) == 1 {
		fmt.Fprintf(writer, "%v\n", 1)
		return
	}

	dp1 := make([]int, len(s))
	dp2 := make([]int, len(s))

	dp1[0] = 1
	if s[0] != s[1] {
		dp1[1] = 2
	}
	dp2[1] = 1
	for i := 2; i < len(s); i++ {

		// dp1
		if s[i] == s[i-1] {
			dp1[i] = max(dp1[i], dp2[i-1]) + 1
		} else {
			dp1[i] = max(dp1[i], dp1[i-1], dp2[i-1]) + 1
		}

		// dp2
		if s[i-2] == s[i-1] {
			dp2[i] = max(dp2[i], dp1[i-2]) + 1
		} else {
			dp2[i] = max(dp2[i], dp1[i-2], dp2[i-2]) + 1
		}
	}
	//fmt.Fprintf(writer, "%v\n", dp1)
	fmt.Fprintf(writer, "%v\n", max(dp1[len(s)-1], dp2[len(s)-1]))
}
