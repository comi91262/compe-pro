package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 1_000_000_000 + 7

func pow(a, x int) int {
	r := 1
	for x > 0 {
		if x&1 == 1 {
			r = r * a % mod
		}
		a = a * a % mod
		x >>= 1
	}
	return r
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	sort.Ints(a)

	valid := true
	if n%2 == 1 {
		cnt := 0
		for i := 0; i < (n+1)/2; i++ {
			if i == 0 {
				if a[i] != cnt {
					valid = false
					break
				}
				cnt += 2
				continue
			}

			if a[2*i-1] != cnt || a[2*i] != cnt {
				valid = false
				break
			}
			cnt += 2
		}
	} else {
		cnt := 1
		for i := 0; i < (n+1)/2; i++ {
			if a[2*i] != cnt || a[2*i+1] != cnt {
				valid = false
				break
			}
			cnt += 2
		}
	}
	if !valid {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}

	fmt.Fprintf(writer, "%v\n", pow(2, n/2)%mod)
}
