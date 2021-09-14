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
	var s = strings.Split("DiscoPresentsDiscoveryChannelProgrammingContest2016", "")
	var w int
	fmt.Fscan(reader, &w)

	ans := []string{}
	stack := ""
	for i := 0; i < len(s); i++ {
		stack += s[i]
		if (i+1)%w == 0 {
			ans = append(ans, stack)
			stack = ""
			continue
		}
	}
	if stack != "" {
		ans = append(ans, stack)
	}

	for i := range ans {
		fmt.Fprintf(writer, "%v\n", ans[i])
	}
}
