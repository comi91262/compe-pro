package algo

import "math"

func isSquare(n int) bool {
	pre := int(math.Sqrt(float64(n)))

	return n == pre*pre
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

// x = y のとき 一番左のものを返す (std::min)
func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func mod(x, y int) int {
	x = x % y
	if x >= 0 {
		return x
	}
	if y < 0 {
		return x - y
	}
	return x + y
}

func powMod(a, x, d int) int {
	var r int = 1
	for x > 0 {
		if x&1 == 1 {
			r = r * a % d
		}
		a = a * a % d
		x >>= 1
	}
	return r
}

func pow(a, x int) int {
	r := 1
	for x > 0 {
		if x&1 == 1 {
			r *= a
		}
		a *= a
		x >>= 1
	}
	return r
}

func floorLog(x int) int {
	if x == 1 {
		return 0
	}
	r := 1
	prod := 2
	for x > prod {
		prod *= 2
		r++
	}

	return r
}

func ceilLog(x int) int {
	if x == 1 {
		return 0
	}
	r := 1
	prod := 2
	for x > prod {
		prod *= 2
		r++
	}

	if x < prod {
		return r - 1
	} else {
		return r
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	d := gcd(a, b)
	return a / d * b
}

func fact(n int) int {
	r := 1
	for n > 0 {
		r *= n
		n--
	}
	return r
}

func factMod(n int, p int) int {
	r := 1
	for n > 0 {
		r = r * n % p
		n--
	}
	return r
}

func permutation(n int, k int) int {
	if n < k {
		return 0
	}

	prod := 1
	for k > 0 {
		prod *= n - k + 1
		k--
	}
	return prod
}

func combination(n, k int) int {
	r := 1
	for d := 1; d <= k; d++ {
		r *= n
		n--
		r /= d
	}

	return r
}

func makePrimes(n int) []int {
	r := []int{}
	prime := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		prime[i] = 1
	}

	prime[0], prime[1] = 0, 0
	for p := 2; p*p < n; p++ {
		if prime[p] > 0 {
			for x := 0; p*(x+2) <= n; x++ {
				prime[p*(x+2)] = 0
			}
		}
	}
	for p := 2; p < n+1; p++ {
		if prime[p] > 0 {
			r = append(r, p)
		}
	}

	return r
}

// 約数列挙
func divisor(n int) []int {
	var r []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			r = append(r, i)
			if i*i != n {
				r = append(r, n/i)
			}
		}
	}
	return r
}

// 素因数分解
func getSpf(n int) []int {
	var spf = make([]int, n+1)

	for i := 0; i <= n; i++ {
		spf[i] = i
	}
	for i := 2; i*i <= n; i++ {
		if spf[i] == i {
			for j := i * i; j <= n; j += i {
				if spf[j] == j {
					spf[j] = i
				}
			}
		}
	}

	return spf
}

func primeFactor(n int, spf []int) map[int]int {
	var m = map[int]int{}
	for n != 1 {
		m[spf[n]]++
		n /= spf[n]
	}
	return m
}

func primeFactor2(n int) map[int]int {
	var m = map[int]int{}

	for i := 2; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		for n%i == 0 {
			m[i]++
			n /= i
		}
	}
	if n != 1 {
		m[n]++
	}
	return m
}

// nCk (mod p) の計算の高速化 (k <= n <= 10_000_000)
// nCk = n! / k! / (n-k)! (mod p)
//     = n! * (k!)^(-1) * ((n-k)!)^(-1) (mod p)
//     = fact[n] * factInv[k] * factInv[n-k] (mod p)
//
// fact[i] = i * fact[i-1] (mod p)
// factInv[i] = i!^(-1) = inv[i] * fact_inv[i-1] (mod p, inv[i] = (i)^(-1))
// 前処理　O(n)
// クエリ  O(1)
type FastNck struct {
	fact    []int
	factInv []int
	prime   int
}

func (table *FastNck) Create(n, p int) {
	table.fact = make([]int, n+1)
	table.factInv = make([]int, n+1)
	table.prime = p

	table.fact[0], table.fact[1] = 1, 1
	table.factInv[0], table.factInv[1] = 1, 1

	inv := make([]int, n+1)
	inv[1] = 1
	for i := 2; i < n+1; i++ {
		table.fact[i] = table.fact[i-1] * i % p
		inv[i] = p - inv[p%i]*(p/i)%p
		table.factInv[i] = table.factInv[i-1] * inv[i] % p
	}
}

func (table *FastNck) Combination(n, k int) int {
	if n < k {
		return 0
	}
	if k == 0 {
		return 1
	}

	return table.fact[n] * (table.factInv[k] * table.factInv[n-k] % table.prime) % table.prime
}

// k <= 10_000_000 <=  n <= 1_000_000_000
type FastNck2 struct {
	factInv []int
	prime   int
}

func (table *FastNck2) Create(k, p int) {
	table.factInv = make([]int, k+1)
	table.prime = p
	inv := make([]int, k+1)

	table.factInv[0], table.factInv[1] = 1, 1
	inv[1] = 1

	for i := 2; i < k; i++ {
		inv[i] = p - inv[p%i]*(p/i)%p
		table.factInv[i] = table.factInv[i-1] * inv[i] % p
	}
}

func (table *FastNck2) Combination(n, k int) int {
	r := 1
	for i := n; i >= n-k+1; i-- {
		r *= i
		r %= table.prime
	}
	return r * table.factInv[k] % table.prime
}
