package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, p, q int

	fmt.Fscan(reader, &n, &p, &q)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	//	var ba = make([](*big.Int), n)
	//	for i := 0; i < n; i++ {
	//		ba[i] = big.NewInt(int64(a[i]))
	//	}

	sum := 0
	//	bp := big.NewInt(int64(p))
	//	bq := big.NewInt(int64(q))
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				for l := k + 1; l < n; l++ {
					for m := l + 1; m < n; m++ {
						//	tmp := big.NewInt(1)
						//	tmp.Mul(tmp, ba[i])
						//	tmp.Mul(tmp, ba[j])
						//	tmp.Mod(tmp, bp)
						//	tmp.Mul(tmp, ba[k])
						//	tmp.Mod(tmp, bp)
						//	tmp.Mul(tmp, ba[l])
						//	tmp.Mod(tmp, bp)
						//	tmp.Mul(tmp, ba[m])
						//	tmp.Mod(tmp, bp)
						//		if tmp.Cmp(bq) == 0 {
						//			sum++
						//		}
						if a[i]%p*a[j]%p*a[k]%p*a[l]%p*a[m]%p == q {
							sum++
						}
					}
				}
			}
		}
	}
	fmt.Fprintf(writer, "%d\n", sum)

}
