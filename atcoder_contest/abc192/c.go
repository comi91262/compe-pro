package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func g1(x int) int {
	s := strconv.Itoa(x)
	ss := strings.Split(s, "")
	sort.Slice(ss, func(i, j int) bool { return ss[i] > ss[j] })
	s = strings.Join(ss, "")
	n, _ := strconv.Atoi(s)
	return n
}
func g2(x int) int {
	s := strconv.Itoa(x)
	ss := strings.Split(s, "")
	sort.Slice(ss, func(i, j int) bool { return ss[i] < ss[j] })
	s = strings.Join(ss, "")
	n, _ := strconv.Atoi(s)
	return n
}
func f(x int) int {
	return g1(x) - g2(x)
}
func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)

	a := n
	for i := 0; i < k; i++ {
		a = f(a)
	}
	fmt.Fprintf(writer, "%v\n", a)
}
