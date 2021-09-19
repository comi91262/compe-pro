package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func swap(a []string, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func reverse(a []string, i, j int) {
	for k := 0; k < (j-i+1)/2; k++ {
		swap(a, i+k, j-k)
	}
}
func nextPermutation(a []string) {
	var lastIndex = len(a) - 1

	if lastIndex < 1 {
		return
	}

	var i = lastIndex - 1
	for i >= 0 && a[i] >= a[i+1] {
		i -= 1
	}
	if i < 0 {
		reverse(a, 0, lastIndex)
		return
	}

	var j = lastIndex
	for j > i+1 && a[j] <= a[i] {
		j -= 1
	}

	swap(a, i, j)
	reverse(a, i+1, lastIndex)
}

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	t := make([]string, n)
	copy(t, s)

	for i := 0; i < fact(n); i++ {
		isOk := true
		if strings.Join(s, "") == strings.Join(t, "") {
			isOk = false
		}
		for j := 0; j < n/2; j++ {
			t[j], t[n-j-1] = t[n-j-1], t[j]
		}
		if strings.Join(s, "") == strings.Join(t, "") {
			isOk = false
		}
		for j := 0; j < n/2; j++ {
			t[j], t[n-j-1] = t[n-j-1], t[j]
		}
		if isOk {
			fmt.Fprintf(writer, "%v\n", strings.Join(t, ""))
			return
		}

		nextPermutation(t)
	}
	fmt.Fprintf(writer, "%v\n", "None")
}
