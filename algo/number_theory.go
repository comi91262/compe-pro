package algo

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

func combination(n int, k int) int {
	return permutation(n, k) / permutation(k, k)
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

// k <= n <= 10_000_000
func initFactTb(n, p int) ([]int, []int, []int) {
	fact := make([]int, n+1)
	factInv := make([]int, n+1)
	inv := make([]int, n+1)
	fact[1] = 1
	factInv[1] = 1
	inv[1] = 1
	for i := 2; i < n+1; i++ {
		fact[i] = fact[i-1] * i % p
		inv[i] = p - inv[p%i]*(p/i)%p
		factInv[i] = factInv[i-1] * inv[i] % p
	}
	return fact, factInv, inv
}

func combinationMod(n, k, p int, fact, factInv, inv []int) int {
	return fact[n] * (factInv[k] * factInv[n-k] % p) % p
}

// k <= 10_000_000 <=  n <= 1_000_000_000
func initFactTb2(n, p int) []int {
	factInv := make([]int, n+1)
	inv := make([]int, n+1)

	factInv[1] = 1
	inv[1] = 1
	for i := 2; i < n; i++ {
		inv[i] = p - inv[p%i]*(p/i)%p
		factInv[i] = factInv[i-1] * inv[i] % p
	}

	return factInv
}

func combinationMod2(n, k, p int, factInv []int) int {
	ans := 1
	for i := n; i >= n-k+1; i-- {
		ans *= i
		ans %= p
	}
	return ans * factInv[k] % p
}
