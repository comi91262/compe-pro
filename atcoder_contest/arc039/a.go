package main

import (
	"bufio"
	"fmt"
	"os"
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
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	an := []int{a + (9-a/10/10%10)*100, a + (9-a/10%10)*10, a + (9 - a%10)}
	bn := []int{b - (b/10/10%10-1)*100, b - (b/10%10)*10, b - (b % 10)}

	//fmt.Fprintf(writer, "%v %v\n", a, an)
	//fmt.Fprintf(writer, "%v %v\n", b, bn)

	ans := -1 << 60
	for i := 0; i < len(an); i++ {
		for j := 0; j < len(bn); j++ {
			ans = max(ans, a-bn[j], an[i]-b)
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
