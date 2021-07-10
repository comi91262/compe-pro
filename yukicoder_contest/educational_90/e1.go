package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func toDigits(x, base int) []int {
	if x == 0 {
		return []int{0}
	}

	ans := make([]int, 0)
	for x != 0 {
		ans = append(ans, x%base)
		x = x / base
	}
	return ans
}

func main() {
	defer writer.Flush()

	var p, q, r, k int
	fmt.Fscan(reader, &p, &q, &r, &k)

	a := make([]int, 1000000)

	cnt := 0
	m2 := 0
	for i := 1; i <= k; i++ {
		if i == 1 {
			a[i] = p % 10
			continue
		}
		if i == 2 {
			a[i] = q % 10
			continue
		}
		if i == 3 {
			a[i] = r % 10
			continue
		}

		if a[i-1] == r%10 && a[i-2] == q%10 && a[i-3] == p%10 {
			cnt++
			if cnt == 2 {
				m2 = i - 4
				break
			}
		}
		a[i] = (a[i-1] + a[i-2] + a[i-3]) % 10
	}
	if m2 == 0 {
		fmt.Fprintf(writer, "%v", a[k])
		return
	}

	n2 := k % m2

	a = make([]int, 1000000)
	for i := 1; i <= n2; i++ {
		if i == 1 {
			a[i] = p % 10
			continue
		}
		if i == 2 {
			a[i] = q % 10
			continue
		}
		if i == 3 {
			a[i] = r % 10
			continue
		}

		a[i] = (a[i-1] + a[i-2] + a[i-3]) % 10
	}

	fmt.Fprintf(writer, "%v", a[n2])
}
