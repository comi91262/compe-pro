package main
 
import (
	"bufio"
	"fmt"
	"os"
	"sort"
)
 
var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)
 
func main() {
	defer writer.Flush()
 
	var n int
	fmt.Fscan(reader, &n)
	var l = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &l[i])
	}
	sort.Ints(l)
 
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if l[i]+l[j] > l[k] {
					ans++
				}
			}
		}
	}
 
	fmt.Fprintf(writer, "%v\n", ans)
}