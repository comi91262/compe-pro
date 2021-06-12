package main

import (
	"bufio"
	"fmt"
	"os"
)

var dp [100001]int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var k int
	fmt.Fscan(reader, &k)

	const d = 1_000_000_000 + 7

	if k%9 == 0 {
		dp[1] = 1
		for i := 2; i < 10; i++ {
			dp[i] = 2 * dp[i-1]
		}

		if 10 < k {
			for i := 10; i < k+1; i++ {
				sum := 0
				for j := 1; j < 10; j++ {
					sum += dp[i-j]
				}
				dp[i] = sum % d
			}
		}
		fmt.Fprintf(writer, "%d\n", dp[k])
	} else {
		fmt.Fprintf(writer, "%d\n", 0)
	}
}
